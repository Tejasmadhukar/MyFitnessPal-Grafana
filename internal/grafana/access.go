package grafana

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func SendGrafanaRequest(method string, endpoint string, body *strings.Reader) (map[string]interface{}, error) {
	var responseJson map[string]interface{}

	RequestUrl, err := url.JoinPath(config.GRAFANA_HOST, endpoint)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(method, RequestUrl, body)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.GRAFANA_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&responseJson)

	if resp.Status != "200 OK" {
		log.Println(responseJson)
		log.Println(resp.Status)
		return nil, errors.New(responseJson["message"].(string))
	}

	return responseJson, nil
}
