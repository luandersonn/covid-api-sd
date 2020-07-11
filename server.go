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

	// handle /all
	go func() {
		http.HandleFunc("/cities", casesPerCitiesHandler)
		log.Printf("Server \"/cities\" starting on port %v\n", port)
		http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
		wait.Done()
	}()

	// wait goroutines
	wait.Wait()
}

func printRequestInfo(request *http.Request) {
	fmt.Println("New request")
	fmt.Printf("\tURL: \"%v\"\n", request.RequestURI)
	fmt.Printf("\tMethod: %v\n", request.Method)
	fmt.Printf("\tUser-Agent: %v\n", request.Header.Get("User-Agent"))
}

func casesPerCitiesHandler(responceWriter http.ResponseWriter, request *http.Request) {
	printRequestInfo(request)

	data, err := csv.ReadFile("/home/luandersonn/Downloads/casos_coronavirus.csv")
	ensureSuccessStatus(err)

	responseData := util.GetCasesPerCity(data)
	dataJSON, err := json.Marshal(responseData)
	ensureSuccessStatus(err)

	responceWriter.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(responceWriter, string(dataJSON))
	/*
		// get  queries
		for key, value := range request.URL.Query() {
			fmt.Fprintf(responceWriter, "%v = %v\n", key, value)
			b = true
		}
		if b {
			return
		}*/
}

func ensureSuccessStatus(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
