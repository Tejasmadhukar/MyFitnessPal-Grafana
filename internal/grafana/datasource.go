package grafana

import (
	"fmt"
	"net/url"
	"path/filepath"
	"strings"

	"github.com/Tejasmadhukar/MyFitnessPal-Grafana/pkg/config"
)

func AddDataSource(fileID string) error {
	fileurl, err := url.JoinPath(config.GRAFANA_HOST, filepath.Join("static", fileID+".csv"))
	if err != nil {
		return err
	}

	datasourceModel := strings.NewReader(fmt.Sprintf(`{
      "access": "proxy",
      "basicAuth": false,
      "basicAuthUser": "",
      "database": "",
      "isDefault": false,
      "jsonData": {
        "storage": "http"
      },
      "name": "%v",
      "type": "marcusolsson-csv-datasource",
      "uid": "%v",
      "url": "%v",
      "user": "",
      "withCredentials": false
    }`, fileID, fileID, fileurl))

	DataSourceEndpoint := "/api/datasources"

	if _, err = SendGrafanaRequest("POST", DataSourceEndpoint, datasourceModel); err != nil {
		return err
	}

	return nil
}
