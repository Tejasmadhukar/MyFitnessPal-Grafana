package grafana

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func CreateDashboard(fileId string) (string, error) {
	var responseJson map[string]interface{}
	DashboardBody := strings.NewReader(fmt.Sprintf(`{
    "dashboard": %v,
    "isFolder": false,
    "message": "created dashboard for user",
    "overwrite": true,
    "userId": 0
  }`, GetDashboard(fileId)))

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+"/api/dashboards/db", DashboardBody)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.GRAFANA_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseJson)

	if resp.Status != "200 OK" {
		log.Println(responseJson)
		log.Println(resp.Status)
		return "", errors.New(responseJson["message"].(string))
	}

	return responseJson["uid"].(string), nil
}

func CreatePublicDashboard(dashboardId string) (string, error) {
	var responseJson map[string]interface{}
	PublicDashboardBody := strings.NewReader(fmt.Sprintf(`{
    "uid": "%v",
    "isEnabled": true,
    "annotationsEnabled": true
  }`, dashboardId))

	PublicDashboardEndpoint := fmt.Sprintf(`/api/dashboards/uid/%v/public-dashboards`, dashboardId)

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+PublicDashboardEndpoint, PublicDashboardBody)
	if err != nil {
		return "", err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.GRAFANA_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseJson)

	if resp.Status != "200 OK" {
		log.Println(responseJson)
		log.Println(resp.Status)
		return "", errors.New(responseJson["message"].(string))
	}

	return responseJson["accessToken"].(string), nil
}
