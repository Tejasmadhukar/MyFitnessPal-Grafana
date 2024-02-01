package handler

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/grafana"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/go-chi/chi/v5"
)

func Visualize(w http.ResponseWriter, r *http.Request) {
	filename := chi.URLParam(r, "filename") + ".csv"

	_, err := os.Stat(config.ASSETS_DIR + "data/" + filename)
	if errors.Is(err, os.ErrNotExist) {
		log.Println(err)
		newError := models.HtmlClientError{
			Status:       404,
			ErrorMessage: "File that you want to Visualize does not exist.",
		}
		newError.Send(&w)
	} else if err != nil {
		log.Printf("This should not happen %v", err)
		w.WriteHeader(500)
		return
	}

	err = grafana.AddDataSource(filename)
	if err != nil {
		newError := models.HtmlClientError{
			Status:       500,
			ErrorMessage: "Grafana api did not respond successfully Error: \n" + err.Error(),
		}
		newError.Send(&w)
		return
	}

	w.Write([]byte("successfully created datasource now make dashboard"))
}
