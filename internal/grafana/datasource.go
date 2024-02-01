package grafana

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
	"github.com/matoous/go-nanoid/v2"
)

func AddDataSource(fileName string) error {
	fileurl := config.HOST_URL + "/static/data/" + fileName
	newID, err := gonanoid.New()
	if err != nil {
		log.Fatal(err)
	}

	datasourceModel := fmt.Sprintf(`{
      "access": "string",
      "basicAuth": true,
      "basicAuthUser": "string",
      "database": "string",
      "isDefault": false,
      "jsonData": {
        "storage": "http"
      },
      "name": "%v",
      "type": "marcusolsson-csv-datasource",
      "uid": "%v",
      "url": "%v",
      "user": "",
      "withCredentials": true
    }`, fileName, newID, fileurl)

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+"/api/datasources", bytes.NewBuffer([]byte(datasourceModel)))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+config.GRAFANA_TOKEN)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.Status != "200 OK" {
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Printf("Error reading response body %v", err)
			return err
		}
		log.Println(datasourceModel)
		log.Println(string(body))
		log.Println(resp.Status)
		return errors.New("Datasource api did not respond with 200 \n" + string(body))
	}

	return nil
}
