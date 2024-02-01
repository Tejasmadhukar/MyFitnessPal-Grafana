package handler

import (
	"html/template"
	"log"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles(config.ASSETS_DIR + "templates/index.html")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Html template could not be parsed"))
		return
	}
	tmpl.Execute(w, nil)
}
