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
		var NewError models.HtmlClientError
		NewError.Status = 400
		NewError.ErrorMessage = "No file or bad file was sent. Refresh to try again"
		NewError.Send(&w)
		return
	}
	defer file.Close()

	fileName := strings.Split(header.Filename, ".")

	if fileName[1] != "csv" {
		var NewError models.HtmlClientError
		NewError.Status = 400
		NewError.ErrorMessage = "Server only accepts a csv file. Refresh to try again"
		NewError.Send(&w)
		return
	}

	csvReader := csv.NewReader(io.Reader(file))

	csvData, err := csvReader.ReadAll()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading csv file " + err.Error()))
		return
	}

	newFilename := strings.ReplaceAll(fileName[0], "/", "_")
	filepath := config.ASSETS_DIR + "data/" + newFilename + ".csv"

	valid := models.CheckCsvHeaders(csvData[0])
	if !valid {
		var NewError models.HtmlClientError
		NewError.ErrorMessage = "The csv file you sent is not valid. Missing (Date, Calories, Meal), If you have these then please rename them in your csv file. Refresh to try again"
		NewError.Status = 400
		NewError.Send(&w)
		return
	}

	csvData = models.TransformCsv(csvData)

	fileToWrite, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not save file " + err.Error()))
		return
	}

	csvWriter := csv.NewWriter(fileToWrite)
	csvWriter.WriteAll(csvData)

	if err = csvWriter.Error(); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not save file " + err.Error()))
		return
	}

	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/success_validation.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Html template could not be parse"))
		return
	}

	response := successfulResponse{
		Filename: newFilename,
	}

	tmpl.Execute(w, response)

}
