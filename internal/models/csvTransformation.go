package models

import (
	"fmt"
	"strconv"
)

func TransformCsv(csvData [][]string) [][]string {
	DateIndex := -1
	GroupedCsvMap := make(map[string][]string)

	headers := csvData[0]
	for idx, val := range headers {
		if val == "Date" {
			DateIndex = idx
		}
	}

	// litreally doing df.groupBy('Date').sum()
	for i := 1; i < len(csvData); i++ {
		val, ok := GroupedCsvMap[csvData[i][DateIndex]]

		if !ok {
			GroupedCsvMap[csvData[i][DateIndex]] = csvData[i]
			continue
		}

		for idx, initialData := range val {
			if idx == DateIndex {
				continue
			}

			initialVal, err := strconv.ParseFloat(initialData, 32)
			if err != nil {
				continue
			}

			valToAdd, err := strconv.ParseFloat(csvData[i][idx], 32)
			if err != nil {
				continue
			}

			sum := initialVal + valToAdd

			GroupedCsvMap[csvData[i][DateIndex]][idx] = fmt.Sprintf("%.1f", sum)
		}
	}

	TransformedData := make([][]string, len(GroupedCsvMap))

	i := 0
	for _, val := range GroupedCsvMap {
		TransformedData[i] = val
		i++
	}

	return TransformedData
}
