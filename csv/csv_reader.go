package csv

import (
	"encoding/csv"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

var data []CovidData

// CovidData
type CovidData struct {
	PacientCode     string
	PacientAge      string
	PacientGender   string
	PacientDistrict string
	PacientCity     string
	PacientState    string
}

// ReadFile faz o que o nome diz, lê o arquivo
func ReadFile(path string) ([]CovidData, error) {
	// Open file
	var err error
	err = nil
	if data == nil {
		file, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		defer file.Close()

		// Read bytes
		bytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		data, err = parseData(string(bytes))
	}
	return data, err
}

func parseData(data string) ([]CovidData, error) {
	reader := csv.NewReader(strings.NewReader(data))
	cases := []CovidData{}
	c := 0
	for {
		record, err := reader.Read()
		c++
		if c == 1 {
			continue
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		p := CovidData{
			PacientDistrict: record[0],
			PacientCode:     record[3],
			PacientAge:      record[27],
			PacientCity:     record[29],
			PacientState:    record[25],
			PacientGender:   record[33]}
		cases = append(cases, p)
	}
	return cases, nil
}
