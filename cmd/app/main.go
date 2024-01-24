package main

import (
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/router"
	"github.com/go-chi/chi/v5"
)

var (
	r = chi.NewRouter()
)

func init() {
	router.AddRoutes(r)
}

func main() {
	fs := http.FileServer(http.Dir("internal/assets/"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":80", r)
}
