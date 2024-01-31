package models

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

type HtmlClientError struct {
	Status       int32
	ErrorMessage string
}

func (ClientError HtmlClientError) Send(w *http.ResponseWriter) {
	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/error.html")
	if err != nil {
		log.Fatal("Should not happen")
	}

	tmpl.Execute(*w, ClientError)
}
