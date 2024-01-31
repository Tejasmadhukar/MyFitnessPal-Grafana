package models

import "slices"

var (
	ValidHeaders = []string{"Date", "Meal", "Calories"}
)

func CheckCsvHeaders(Headers []string) bool {
	for _, val := range ValidHeaders {
		if !slices.Contains(Headers, val) {
			return false
		}
	}
	return true
}
