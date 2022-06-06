package etl

func discoverCreditTransformer(line []string) []string {
	outputData := []string{
		line[1], line[3], "", line[2],
	}

	return outputData
}

func discoverBankTransformer(line []string) []string {

	outputData := []string{
		line[0], line[1], "", line[3],
	}

	return outputData
}

func mtbTransformer(line []string) []string {

	outputData := []string{
		line[1], line[3], "", line[2],
	}
	return outputData
}
