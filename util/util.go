package util

import (
	"time"

	"github.com/luandersonn/covid-api-sd/csv"
)

// CasesPerCityResponse é a estrutura de resposta para solicitações de cidades e a quantidade de casos
type CasesPerCityResponse struct {
	Date   time.Time      `json:"date"`
	Cities []CasesPerCity `json:"cities"`
}

// CasesPerCity armazena o nome da cidade e o número de casos
type CasesPerCity struct {
	City       string `json:"city"`
	CityCode   string `json:"city_code,omitempty"`
	CasesCount int    `json:"cases_count"`
}

// CityResponse é a estrutura de resposta para solicitações de todos os pacientes de uma cidade específica
type CityResponse struct {
	City       string      `json:"city,omitempty"`
	State      string      `json:"state,omitempty"`
	Code       string      `json:"code,omitempty"`
	CasesCount int         `json:"cases_count,omitempty"`
	Cases      []CovidCase `json:"cases,omitempty"`
}

// CovidCase representa um caso de covid-19
type CovidCase struct {
	Code     string     `json:"code,omitempty"`
	Age      string     `json:"age,omitempty"`
	Gender   string     `json:"gender,omitempty"`
	District string     `json:"district,omitempty"`
	City     string     `json:"city,omitempty"`
	CityCode string     `json:"city_code,omitempty"`
	State    string     `json:"state,omitempty"`
	Date     *time.Time `json:"date,,omitempty"`
}

// Find procura por um caso de covid baseado no comparador passado
// através de uma função lambda
func Find(slice []csv.CovidData, comparer func(csv.CovidData) bool) *csv.CovidData {
	for _, item := range slice {
		if comparer(item) {
			return &item
		}
	}
	return nil
}

// Map seleciona itens de um slice baseado no comparador passado
// através de uma função lambda
func Map(slice []csv.CovidData, comparer func(csv.CovidData) bool) []csv.CovidData {
	result := []csv.CovidData{}
	for _, item := range slice {
		if comparer(item) {
			result = append(result, item)
		}
	}
	return result
}

// GroupBy agrupa
func GroupBy(slice []csv.CovidData, keySelector func(csv.CovidData) string) map[string][]csv.CovidData {
	// Cria um dicionário [chave] -> slice de csv.CovidData
	groups := make(map[string][]csv.CovidData)

	for _, item := range slice {
		groups[keySelector(item)] = append(groups[keySelector(item)], item)
	}
	return groups
}

// Unique remove strings repetidas
func Unique(slice []string) []string {
	keys := make(map[string]bool)
	newList := []string{}
	for _, entry := range slice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newList = append(newList, entry)
		}
	}
	return newList
}
