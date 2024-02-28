package handler

import (
	"errors"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/grafana"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/go-chi/chi/v5"
)

func Visualize(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename") + ".csv"

	_, err := os.Stat(filepath.Join(config.ASSETS_DIR, filename))
	if errors.Is(err, os.ErrNotExist) {
		newError := models.ErrorResponse{
			Status:       404,
			ErrorMessage: "File that you want to Visualize does not exist.",
		}
		newError.SendError(&w)
		return
	} else if err != nil {
		log.Printf("This should not happen %v", err)
		w.WriteHeader(500)
		return
	}

	err = grafana.AddDataSource(filename)
	if err != nil {
		models.SendInternalServerError(&w, "Grafana api did not respond successfully Error: \n"+err.Error())
		return
	}

	w.Write([]byte("successfully created datasource now make dashboard"))

	go func() {
		os.Remove(filepath.Join(config.ASSETS_DIR, filename))
	}()
}
