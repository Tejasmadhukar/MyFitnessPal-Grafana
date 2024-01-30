package handler

import (
	"encoding/csv"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
)

type successfulResponse struct {
	Filename string
}

func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(60 << 20)

	file, header, err := r.FormFile("file")
	if err != nil {
		var NewError models.HtmlClientError
		NewError.Status = 400
		NewError.ErrorMessage = "No file or bad file was sent. Refresh to try again"
		SendError(NewError, &w)
		return
	}
	defer file.Close()

	fileName := strings.Split(header.Filename, ".")

	if fileName[1] != "csv" {
		var NewError models.HtmlClientError
		NewError.Status = 400
		NewError.ErrorMessage = "Server only accepts a csv file. Refresh to try again"
		SendError(NewError, &w)
		return
	}

	csvReader := csv.NewReader(io.Reader(file))

	record, err := csvReader.Read()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading csv file " + err.Error()))
		return
	}

	valid := models.CheckCsvHeaders(record)

	if !valid {
		var NewError models.HtmlClientError
		NewError.ErrorMessage = "The csv file you sent is not valid. Missing (Date, Calories, Meal), If you have these then please rename them in your csv file. Refresh to try again"
		NewError.Status = 400
		SendError(NewError, &w)
		return
	}

	data, err := io.ReadAll(file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error reading file " + err.Error()))
		return
	}

	newFilename := strings.ReplaceAll(fileName[0], "/", "_")

	err = os.WriteFile("internal/assets/data/"+newFilename+".csv", data, 0644)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Could not save file " + err.Error()))
		return
	}

	tmpl, err := template.ParseFiles("internal/handler/templates/success_validation.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Html template could not be parse"))
		return
	}

	response := successfulResponse{
		Filename: newFilename,
	}

	tmpl.Execute(w, response)

}
