package models

import (
	"fmt"
	"slices"
	"strconv"
	"time"
)

type ByDate []time.Time

func (t ByDate) Len() int           { return len(t) }
func (t ByDate) Swap(i, j int)      { t[i], t[j] = t[j], t[i] }
func (t ByDate) Less(i, j int) bool { return t[i].Before(t[j]) }

func TransformCsv(csvData [][]string) [][]string {
	DateIndex := -1
	const shortForm = "2006-01-31"

	headers := csvData[0]
	for idx, val := range headers {
		if val == "Date" {
			DateIndex = idx
		}
	}

	// litreally doing df.groupBy('Date').sum()
	i := 2
	for i < len(csvData) {
		if csvData[i][DateIndex] == csvData[i-1][DateIndex] {
			startIndex := i - 1
			for i < len(csvData) && csvData[i][DateIndex] == csvData[i-1][DateIndex] {
				for idx, val := range csvData[i] {
					if idx == DateIndex {
						continue
					}
					initialVal, err := strconv.ParseFloat(csvData[startIndex][idx], 32)
					if err != nil {
						continue
					}

					valToAdd, err := strconv.ParseFloat(val, 32)
					if err != nil {
						continue
					}

					sum := initialVal + valToAdd

					csvData[startIndex][idx] = fmt.Sprintf("%.1f", sum)
				}
				i++
			}
			csvData = slices.Delete(csvData, startIndex+1, i)
			continue
		}
		i++
	}

	return csvData
}
