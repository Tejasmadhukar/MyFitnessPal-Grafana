package handler

import (
	"encoding/csv"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

type successfulResponse struct {
	Filename string
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(60 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		log.Println(err)
		models.SendBadRequest(&w, `No file or bad file was sent. Refresh to try again`)
		return
	}
	defer file.Close()

	fileName := strings.Split(header.Filename, ".")

	if fileName[1] != "csv" {
		models.SendBadRequest(&w, `Server only accepts a csv file. Refresh to try again`)
		return
	}

	csvReader := csv.NewReader(io.Reader(file))

	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "Error reading csv file "+err.Error())
		return
	}

	newFilename := strings.ReplaceAll(fileName[0], "/", "_")
	filepath := config.ASSETS_DIR + "data/" + newFilename + ".csv"

	err = models.CheckCsvHeaders(csvData[0])
	if err != nil {
		models.SendInternalServerError(&w, err.Error())
		return
	}

	err = models.CheckCsvData(csvData)
	if err != nil {
		models.SendBadRequest(&w, err.Error())
		return
	}

	csvData = models.TransformCsv(csvData)

	fileToWrite, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "Could not create file "+err.Error())
		return
	}

	csvWriter := csv.NewWriter(fileToWrite)
	csvWriter.WriteAll(csvData)

	if err = csvWriter.Error(); err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "Could not write to file "+err.Error())
		return
	}

	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/success_validation.html")
	if err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "HTML template couldn't be parsed")
		return
	}

	response := successfulResponse{
		Filename: newFilename,
	}

	tmpl.Execute(w, response)

}
