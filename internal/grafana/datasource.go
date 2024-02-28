package grafana

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func AddDataSource(fileName string) error {
	host := config.HOST_URL + ":" + config.PORT
	fileurl := host + "/" + filepath.Join("static", fileName)

	datasourceModel := strings.NewReader(fmt.Sprintf(`{
      "access": "proxy",
      "basicAuth": true,
      "basicAuthUser": "basicuser",
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
    }`, fileName, fileName, fileurl))

	req, err := http.NewRequest("POST", config.GRAFANA_HOST+"/api/datasources", datasourceModel)
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
