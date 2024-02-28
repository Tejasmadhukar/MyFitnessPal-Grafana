package handler

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/grafana"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/internal/models"
	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/go-chi/chi/v5"
)

type FinalDashboard struct {
	SnapshotUrl string
}

func Visualize(w http.ResponseWriter, r *http.Request) {
	fileID := chi.URLParam(r, "filename")
	filename := fileID + ".csv"

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

	err = grafana.AddDataSource(fileID)
	if err != nil {
		models.SendInternalServerError(&w, "Grafana datasource api did not respond successfully Error: \n"+err.Error())
		return
	}

	w.Write([]byte("successfully created datasource now make dashboard"))

	snapshotUrl, err := grafana.CreateSnapShot(fileID)
	if err != nil {
		models.SendInternalServerError(&w, "Grafana snapshots api did not respond successfully Error: \n"+err.Error())
	}

	tmpl, err := template.ParseFiles(filepath.Join(config.ASSETS_DIR, "templates", "snapshot_iframe.html"))
	if err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "HTML template couldn't be parsed")
		return
	}

	response := &FinalDashboard{
		SnapshotUrl: snapshotUrl,
	}

	tmpl.Execute(w, response)

	go func() {
		os.Remove(filepath.Join(config.ASSETS_DIR, filename))
	}()
}
