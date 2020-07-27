package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sort"
	"sync"
	"time"

	covid "github.com/luandersonn/covid-api-sd/protofile"
	"github.com/luandersonn/covid-api-sd/util"
	"google.golang.org/grpc"
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

func newRPCRequest() ([]*covid.CovidDataResponse, error) {
	// Faz a ligação para o endereço do servidor RPC de leitura de csv
	rpcAddress := "localhost:5001"
	conn, err := grpc.Dial(rpcAddress, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	// Fecha a conexão ao finalizar o método
	defer conn.Close()

	// Obtem o serviço DataService
	service := covid.NewCovidDataServiceClient(conn)

	// Obtem o contexto com timeout de 15 segundos
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	// Faz a requisição
	stream, err := service.GetDataStream(ctx, &covid.CovidDataRequest{Name: "Server.go"})
	if err != nil {
		return nil, err
	}
	count := 0
	// Faz o stream dos itens
	covidCaseSlice := []*covid.CovidDataResponse{}
	for {
		covidCase, err := stream.Recv()
		count++
		if err == io.EOF {
			return covidCaseSlice, nil
		}
		if err != nil {
			return nil, err
		}
		if count > 10 && count < 20 {
			if covidCase.GetPacientDistrict() == "" {
				fmt.Println("NULL")
			} else {
				fmt.Println(covidCase.GetPacientDistrict())
			}
		}
		covidCaseSlice = append(covidCaseSlice, covidCase)
	}
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

	data, err := newRPCRequest()
	ensureSuccessStatus(err)

	keySelector := func(item *covid.CovidDataResponse) string {
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

	data, err := newRPCRequest()
	ensureSuccessStatus(err)

	comparer := func(item *covid.CovidDataResponse) bool {
		return item.PacientCode == pacientCode
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
			//Date:     result.Date,
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

	data, err := newRPCRequest()
	ensureSuccessStatus(err)

	comparer := func(item *covid.CovidDataResponse) bool {
		return item.CityCode == cityCode
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
				//Date:     covidCase.Date,
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
