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

func CreateSnapShot(fileId string) (string, error) {
	var responseJson map[string]interface{}
	snapShotBody := strings.NewReader(fmt.Sprintf(`{
    "dashboard": %v,
    "expires": 0,
    "external": false,
    "name": "%v"
  }`, GetDashboard(fileId), fileId))

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+"/api/snapshots", snapShotBody)
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
		return "", errors.New(fmt.Sprintf("Datasource api did not respond with 200 \n %+v", responseJson))
	}

	return responseJson["url"].(string), nil
}
