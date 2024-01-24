package handler

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, _ *http.Request) {
	tmpl, err := template.ParseFiles("internal/handler/templates/index.html")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Html template could not be parsed"))
	}
	tmpl.Execute(w, nil)
}
