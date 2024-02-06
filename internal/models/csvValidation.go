package models

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"time"
)

var (
	ValidHeaders = []string{"Date", "Meal", "Calories"}
)

func CheckCsvHeaders(Headers []string) error {
	for _, val := range ValidHeaders {
		if !slices.Contains(Headers, val) {
			return errors.New(`The csv file you sent is not valid. Missing (Date, Calories, Meal), If you have these then please rename them in your csv file. Refresh to try again`)
		}
	}
	return nil
}

func CheckCsvData(csvData [][]string) error {
	DateIndex, CalroiesIndex := -1, -1
	const shortForm = "2006-01-31"

	for idx, val := range csvData[0] {
		if val == "Date" {
			DateIndex = idx
		} else if val == "Calories" {
			CalroiesIndex = idx
		}
	}

	for idx, val := range csvData[1:] {
		_, err := strconv.ParseFloat(val[CalroiesIndex], 32)
		if err != nil {
			errorMessage := fmt.Sprintf(`Your csv file contains is not of valid type. At index %d Calories column could not be parsed into a floating point number. Here is the underlying error message %s`, idx+1, err.Error())
			return errors.New(errorMessage)
		}

		_, err = time.Parse(shortForm, val[DateIndex])
		if err != nil {
			errorMessage := fmt.Sprintf(`Your csv file contains is not of valid type. At index %d Date column could not be parsed into a DateTime object. Here is the underlying error message %s`, idx+1, err.Error())
			return errors.New(errorMessage)
		}
	}

	return nil
}
