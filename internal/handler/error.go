package handler

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func SendError(ClientError models.HtmlClientError, w *http.ResponseWriter) {
	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/error.html")
	if err != nil {
		fmt.Println("should not happen")
		return
	}

	tmpl.Execute(*w, ClientError)
}
