package models

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestTransformCsv(t *testing.T) {
	filepath := "output" + ".csv"

	file, err := os.Open(filepath)
	if err != nil {
		t.Fatal(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(io.Reader(file))

	csvData, err := csvReader.ReadAll()
	if err != nil {
		t.Fatal(err)
	}

	bef := len(csvData)

	t1 := time.Now()
	csvData = TransformCsv(csvData)
	if err != nil {
		t.Fatal(err)
	}
	t2 := time.Now()

	fmt.Println("length before = ", bef, " length after = ", len(csvData))
	fmt.Println("Function took ", t2.Sub(t1))

	fmt.Println("Headers = ", csvData[0])

	DuplicateMap := make(map[string]bool)
	DateIndex := -1

	headers := csvData[0]
	for idx, val := range headers {
		if val == "Date" {
			DateIndex = idx
		}
	}
	for idx, val := range csvData {
		if _, ok := DuplicateMap[val[DateIndex]]; ok {
			t.Fatalf("Found duplicate in the resulting csvData at index %d", idx)
		}
	}

}
