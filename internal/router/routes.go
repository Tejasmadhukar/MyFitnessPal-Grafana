package router

import (
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func AddRoutes(r *chi.Mux) {
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Heartbeat("/health"))
	r.Get("/", handler.Index)
	r.Post("/upload", handler.Upload)
}
