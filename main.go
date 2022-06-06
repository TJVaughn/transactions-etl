package main

import (
	"fmt"
	"os"
	"strings"

	"transactions-etl/etl"
	"transactions-etl/utils"
)

func app() {

	inputs, err := os.ReadDir("inputs/")

	utils.Check(err)

	for _, file := range inputs {

		fmt.Println(file.Name())

		inputFile, err := os.Open("inputs/" + file.Name())

		utils.Check(err)

		if len(strings.Split(strings.ToLower(inputFile.Name()), ".csv")) == 2 { // IF it's a csv

			bank := ""
			name := strings.ToLower(inputFile.Name())

			if strings.Contains(name, "mtb") {
				bank = "mtb"
			}
			if strings.Contains(name, "discover-credit") {
				bank = "dc"
			}
			if strings.Contains(name, "discover-bank") {
				bank = "db"
			}

			if bank == "" {
				panic("Bank not supported")
			}

			etl.CsvETL(inputFile, bank)

		} else {
			panic("File not supported")
		}
	}

	etl.ConsolidateOutputs()

}

func main() {
	fmt.Println("Program starting...")

	app()
}
