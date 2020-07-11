package util

import (
	"sort"
	"time"

	"github.com/luandersonn/covid-api-sd/csv"
)

// CasesPerCityResponse é a estrutura de resposta para solicitações de cidades e a quantidade de casos
type CasesPerCityResponse struct {
	Time   time.Time      `json:"date"`
	Cities []CasesPerCity `json:"cities"`
}

// CasesPerCity armazena o nome da cidade e o número de casos
type CasesPerCity struct {
	City  string `json:"city"`
	Cases int    `json:"cases"`
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
		Time:   time.Now(),
		Cities: mapCities(cities),
	}
	// Ordena as cidades a partir do maior número de casos
	sort.SliceStable(response.Cities, func(i, j int) bool { return response.Cities[i].Cases > response.Cities[j].Cases })

	return response
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
