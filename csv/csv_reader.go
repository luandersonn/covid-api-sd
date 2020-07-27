package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	covid "github.com/luandersonn/covid-api-sd/protofile"
	"github.com/utfutil"
	"google.golang.org/grpc"
)

var data []*covid.CovidDataResponse

const (
	filePath = "/home/luandersonn/Downloads/casos_coronavirus.csv"
)

type csvReader struct{}

func (reader *csvReader) GetData(ctx context.Context, request *covid.CovidDataRequest) (*covid.CovidDataResponse, error) {
	fmt.Printf("New RPC request from %v: GetData(..)\n", request.GetName())
	result, err := readFile(filePath)
	if err == nil {
		return result[0], nil
	}
	return nil, err
}

func (reader *csvReader) GetDataStream(request *covid.CovidDataRequest, stream covid.CovidDataService_GetDataStreamServer) error {
	// Lê o arquivo
	fmt.Printf("New RPC request from %v: GetDataStream(..) \n", request.GetName())
	cases, err := readFile(filePath)
	if err != nil {
		return err
	}
	for _, covidCase := range cases {
		if err = stream.Send(covidCase); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	port := ":5001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error (Listen): %v", err)
	}
	fmt.Printf("CSV service is listen on %v\n", lis.Addr())
	s := grpc.NewServer()
	covid.RegisterCovidDataServiceServer(s, &csvReader{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Error (Serve): %v", err)
	}
}

// ReadFile faz o que o nome diz, lê o arquivo
func readFile(path string) ([]*covid.CovidDataResponse, error) {
	// Open file
	var err error
	err = nil
	if data == nil {

		file, err := utfutil.OpenFile(path, utfutil.UTF8)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		defer file.Close()

		// Read bytes
		bytes, err := utfutil.ReadFile(path, utfutil.UTF8)
		if err != nil {
			return nil, err
		}
		data, err = parseData(string(bytes))
	}
	return data, err
}

func parseData(data string) ([]*covid.CovidDataResponse, error) {
	reader := csv.NewReader(strings.NewReader(data))
	cases := []*covid.CovidDataResponse{}
	c := 0
	for {
		record, err := reader.Read()
		c++
		if c == 1 { // CABELHAÇO
			continue
		}
		if err == io.EOF { // FIM DO ARQUIVO
			break
		}
		if err != nil {
			log.Fatal(err)
			return nil, err
		}

		/*var date *time.Time = nil
		result, err := time.Parse("2006-01-02 15:04:05.0", record[20])
		if err == nil {
			date = &result
		}*/

		p := &covid.CovidDataResponse{
			PacientDistrict: record[0],
			CityCode:        record[2],
			PacientCode:     record[3],
			Date:            record[20],
			PacientAge:      record[27],
			PacientCity:     record[29],
			PacientState:    record[25],
			PacientGender:   record[33]}
		cases = append(cases, p)
	}
	return cases, nil
}
