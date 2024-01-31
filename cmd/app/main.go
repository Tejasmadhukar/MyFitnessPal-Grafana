package main

import (
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/router"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/go-chi/chi/v5"
)

var (
	r = chi.NewRouter()
)

func init() {
	router.AddRoutes(r)
}

func main() {
	fs := http.FileServer(http.Dir(config.ASSETS_DIR))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":"+config.PORT, r)
}
