package etl

import (
	"encoding/csv"
	"os"
	"transactions-etl/utils"
)

func ConsolidateOutputs() {
	outputs, err := os.ReadDir("outputs")
	utils.Check(err)

	allTransactions, err := os.Create("outputs/allTransactions.csv")

	utils.Check(err)

	writer := csv.NewWriter(allTransactions)

	for _, fileName := range outputs {
		filePtr, err := os.Open("outputs/" + fileName.Name())
		utils.Check(err)
		reader := csv.NewReader(filePtr)
		file, err := reader.ReadAll()
		utils.Check(err)
		for _, line := range file {
			writer.Write(line)
		}
	}

}
