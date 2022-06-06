package etl

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
	"transactions-etl/utils"
)

func csvReader(inputFile io.Reader, bank string) ([][]string, string) {

	reader := csv.NewReader(inputFile)

	allRecords, err := reader.ReadAll()

	utils.Check(err)

	return allRecords, bank
}

func csvWriter(values [][]string, file string) {
	outputFile, err := os.Create("outputs/" + file)

	utils.Check(err)

	writer := csv.NewWriter(outputFile)

	err = writer.WriteAll(values)

	utils.Check(err)

	err = outputFile.Close()

	utils.Check(err)

	fmt.Printf("File written. \n")
}

func CsvETL(inputFile *os.File, bank string) {
	allRecords, bank := csvReader(inputFile, bank)

	outputDataValues := make([][]string, 0)

	for _, line := range allRecords {

		var transform []string

		if bank == "dc" {

			transform = discoverCreditTransformer(line)

		} else if bank == "db" {

			transform = discoverBankTransformer(line)

		} else if bank == "mtb" {

			transform = mtbTransformer(line)

		} else {

			panic("bank not supported")

		}
		outputDataValues = append(outputDataValues, transform)
	}

	fmt.Println("Writing to file...")

	parsedFile := strings.ToLower(inputFile.Name())

	parsedFile = strings.Split(parsedFile, "inputs/")[1]

	csvWriter(outputDataValues, parsedFile)

}
