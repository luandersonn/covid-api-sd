package util

import (
	"sort"
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
	City  string `json:"city"`
	Cases int    `json:"cases"`
}

// CovidCase representa um caso de covid-19
type CovidCase struct {
	Code     string     `json:"code,omitempty"`
	Age      string     `json:"age,omitempty"`
	Gender   string     `json:"gender,omitempty"`
	District string     `json:"district,omitempty"`
	City     string     `json:"city,omitempty"`
	State    string     `json:"state,omitempty"`
	Date     *time.Time `json:"date,,omitempty"`
}

// GetCasesPerCity filtra os dados necessários a partir dos dados brutos.
// Neste caso, ele retorna de forma ordenada a lista de cidades junto a sua quandidade de casos
// da cidade com maior quantidade de casos para a menor
func GetCasesPerCity(data []csv.CovidData) CasesPerCityResponse {
	// Dicionário chave - valor (Nome da cidade - número de casos)
	cities := make(map[string]int)
	for _, value := range data {
		cities[value.PacientCity]++
	}

	response := CasesPerCityResponse{
		Date:   time.Now(),
		Cities: mapCities(cities),
	}
	// Ordena as cidades a partir do maior número de casos
	sort.SliceStable(response.Cities, func(i, j int) bool { return response.Cities[i].Cases > response.Cities[j].Cases })

	return response
}

// Find procura por um caso de covid baseado no comparador passado
// através de uma função lambda
func Find(slice []csv.CovidData, comparer func(csv.CovidData) bool) *csv.CovidData {
	for _, n := range slice {
		if comparer(n) {
			return &n
		}
	}
	return nil
}

func mapCities(cities map[string]int) []CasesPerCity {
	data := []CasesPerCity{}
	for key, value := range cities {
		data = append(data, CasesPerCity{City: key, Cases: value})
	}
	return data
}

func unique(list []string) []string {
	keys := make(map[string]bool)
	newList := []string{}
	for _, entry := range list {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			newList = append(newList, entry)
		}
	}
	return newList
}
