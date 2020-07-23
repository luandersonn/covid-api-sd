package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sync"

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

	/*// handle /city
	go func() {
		http.HandleFunc("/pacient", getPacientHandler)
		log.Printf("Server \"/pacient\" starting on port %v\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		wait.Done()
	}()*/

	// wait goroutines
	wait.Wait()
}

func printRequestInfo(request *http.Request) {
	fmt.Println("New request")
	fmt.Printf("\tURL: \"%v\"\n", request.RequestURI)
	fmt.Printf("\tMethod: %v\n", request.Method)
	fmt.Printf("\tUser-Agent: %v\n", request.Header.Get("User-Agent"))
}

func casesPerCitiesHandler(responseWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	data, err := csv.ReadFile("/home/luandersonn/Downloads/casos_coronavirus.csv")
	ensureSuccessStatus(err)

	responseData := util.GetCasesPerCity(data)
	dataJSON, err := json.Marshal(responseData)
	ensureSuccessStatus(err)

	responseWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(responseWriter, string(dataJSON))
}

func getPacientHandler(responseWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	pacientCode := request.URL.Query().Get("id")
	if pacientCode == "" {
		http.Error(responseWriter, "Bad request: Query param: id", http.StatusBadRequest)
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

func ensureSuccessStatus(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
