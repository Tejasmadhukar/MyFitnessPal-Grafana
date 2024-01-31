package grafana

import (
	"bytes"
	"fmt"
	"log"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/matoous/go-nanoid/v2"
)

func AddDataSource(fileName string) error {

	fileurl := config.ASSETS_DIR + "data/" + fileName
	newID, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)
	}

	datasourceModel := fmt.Sprintf(`{
    "access": "string",
    "basicAuth": true,
    "basicAuthUser": "string",
    "database": "string",
    "isDefault": true,
    "jsonData": {},
    "name": "marcusolsson-csv-datasource",
    "secureJsonData": {},
    "type": "marcusolsson-csv-datasource",
    "uid": "%v",
    "url": "%v",
    "user": "string",
    "withCredentials": true
  }`, newID, fileurl)

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+"/api/datasources", bytes.NewBuffer([]byte(datasourceModel)))
	if err != nil {
		log.Fatal("Should not happen")
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer resp.Body.Close()

	if resp.Status != "200" {
		return err
	}

	return nil
}
