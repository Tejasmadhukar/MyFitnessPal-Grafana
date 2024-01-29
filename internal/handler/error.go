package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
)

func SendError(ClientError models.HtmlClientError, w *http.ResponseWriter) {
	tmpl, err := template.ParseFiles("internal/handler/templates/error.html")
	if err != nil {
		fmt.Println("should not happen")
		return
	}

	tmpl.Execute(*w, ClientError)
}
