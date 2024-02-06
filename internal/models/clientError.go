package models

import (
	"html/template"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

type ErrorResponse struct {
	Status       int32
	ErrorMessage string
}

func (Error ErrorResponse) SendError(w *http.ResponseWriter) {
	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/error.html")
	if err != nil {
		panic("should not happen, BAD TEMPLATE   " + err.Error())
	}

	tmpl.Execute(*w, Error)
}

func SendBadRequest(w *http.ResponseWriter, message string) {
	Error := &ErrorResponse{
		Status:       400,
		ErrorMessage: message,
	}

	Error.SendError(w)
}

func SendInternalServerError(w *http.ResponseWriter, message string) {
	Error := &ErrorResponse{
		Status:       500,
		ErrorMessage: message,
	}

	Error.SendError(w)
}
