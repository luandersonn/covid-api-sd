package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	"github.com/luandersonn/covid-api-sd/csv"
	"github.com/luandersonn/covid-api-sd/util"
)

func main() {
	var wait sync.WaitGroup
	wait.Add(3)

	port := 8080

	// handle /cities
	go func() {
		http.HandleFunc("/cities", casesPerCitiesHandler)
		log.Printf("Server \"/cities\" starting on port %v\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		wait.Done()
	}()

	// handle /pacient
	go func() {
		http.HandleFunc("/pacient", getPacientHandler)
		log.Printf("Server \"/pacient\" starting on port %v\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		wait.Done()
	}()

	// handle /city
	go func() {
		http.HandleFunc("/city", getCityHandler)
		log.Printf("Server \"/city\" starting on port %v\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		wait.Done()
	}()

	// wait goroutines
	wait.Wait()
}

// Printa as informações de cada requisição
func printRequestInfo(request *http.Request) {
	fmt.Println("New request")
	fmt.Printf("\tURL: \"%v\"\n", request.RequestURI)
	fmt.Printf("\tMethod: %v\n", request.Method)
	fmt.Printf("\tUser-Agent: %v\n", request.Header.Get("User-Agent"))
}

// Obtém o número de casos por cidade
func casesPerCitiesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	data, err := csv.ReadFile("/home/luandersonn/Downloads/casos_coronavirus.csv")
	ensureSuccessStatus(err)

	keySelector := func(item csv.CovidData) string {
		return item.PacientCity
	}
	responseData := util.CasesPerCityResponse{Date: time.Now()}
	for _, value := range util.GroupBy(data, keySelector) {
		city := util.CasesPerCity{
			City:       value[0].PacientCity,
			CityCode:   value[0].CityCode,
			CasesCount: len(value),
		}
		responseData.Cities = append(responseData.Cities, city)
	}
	// Ordena da cidade com maior quantidades de casos para a menor
	comparer := func(i, j int) bool {
		return responseData.Cities[i].CasesCount > responseData.Cities[j].CasesCount
	}
	sort.SliceStable(responseData.Cities, comparer)

	dataJSON, err := json.Marshal(responseData)
	ensureSuccessStatus(err)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(responseWriter, string(dataJSON))
}

// Obtém detalhes de um paciente através do código do paciente
func getPacientHandler(responseWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	pacientCode := request.URL.Query().Get("code")
	if pacientCode == "" {
		http.Error(responseWriter, "Bad request: Query param: code", http.StatusBadRequest)
		return
	}

	data, err := csv.ReadFile("/home/luandersonn/Downloads/casos_coronavirus.csv")
	ensureSuccessStatus(err)

	comparer := func(x csv.CovidData) bool {
		return x.PacientCode == pacientCode
	}

	result := util.Find(data, comparer)
	responseData := util.CovidCase{}
	if result != nil {
		responseData = util.CovidCase{
			Age:      result.PacientAge,
			Gender:   result.PacientGender,
			District: result.PacientDistrict,
			City:     result.PacientCity,
			CityCode: result.CityCode,
			State:    result.PacientState,
			Code:     result.PacientCode,
			Date:     result.Date,
		}
	}
	dataJSON, err := json.Marshal(responseData)
	ensureSuccessStatus(err)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(responseWriter, string(dataJSON))
}

// Obtém todos os casos de uma cidade pelo código da cidade
func getCityHandler(responseWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	cityCode := request.URL.Query().Get("code")
	if cityCode == "" {
		http.Error(responseWriter, "Bad request: Query param: code", http.StatusBadRequest)
		return
	}

	data, err := csv.ReadFile("/home/luandersonn/Downloads/casos_coronavirus.csv")
	ensureSuccessStatus(err)

	comparer := func(x csv.CovidData) bool {
		return x.CityCode == cityCode
	}

	responseData := util.CityResponse{}
	for _, covidCase := range util.Map(data, comparer) {

		responseData.City = covidCase.PacientCity
		responseData.State = covidCase.PacientState
		responseData.Code = covidCase.CityCode

		responseData.Cases = append(responseData.Cases,
			util.CovidCase{
				Age:      covidCase.PacientAge,
				Gender:   covidCase.PacientGender,
				District: covidCase.PacientDistrict,
				Code:     covidCase.PacientCode,
				Date:     covidCase.Date,
			})
	}

	responseData.CasesCount = len(responseData.Cases)

	dataJSON, err := json.Marshal(responseData)
	ensureSuccessStatus(err)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(responseWriter, string(dataJSON))
}

func ensureSuccessStatus(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
