package handler

import (
	"errors"
	"fmt"
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
	EmbedUrl string
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

	dashboardID, err := grafana.CreateDashboard(fileID)
	if err != nil {
		models.SendInternalServerError(&w, "Grafana dashboard api did not respond successfully Error: \n"+err.Error())
	}

	public_accessToken, err := grafana.CreatePublicDashboard(dashboardID)
	if err != nil {
		models.SendInternalServerError(&w, "Grafana public_dashboard api did not respond successfully Error: \n"+err.Error())
	}

	tmpl, err := template.ParseFiles(filepath.Join(config.ASSETS_DIR, "templates", "snapshot_iframe.html"))
	if err != nil {
		log.Println(err)
		models.SendInternalServerError(&w, "HTML template couldn't be parsed")
		return
	}

	iframeUrl := fmt.Sprintf("%v/public-dashboards/%v", config.GRAFANA_HOST, public_accessToken)

	response := &FinalDashboard{
		EmbedUrl: iframeUrl,
	}

	tmpl.Execute(w, response)

	// go func() {
	// 	os.Remove(filepath.Join(config.ASSETS_DIR, filename))
	// }()
}
