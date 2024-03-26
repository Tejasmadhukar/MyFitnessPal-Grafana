package grafana

import (
	"fmt"
	"strings"
)

func CreateDashboard(fileId string) (string, error) {
	DashboardBody := strings.NewReader(fmt.Sprintf(`{
    "dashboard": %v,
    "isFolder": false,
    "message": "created dashboard for user",
    "overwrite": true,
    "userId": 0
  }`, GetDashboard(fileId)))

	DashBoardEndpoint := "/api/dashboards/db"

	resBody, err := SendGrafanaRequest("POST", DashBoardEndpoint, DashboardBody)
	if err != nil {
		return "", err
	}

	return resBody["uid"].(string), nil
}

func CreatePublicDashboard(dashboardId string) (string, error) {
	PublicDashboardBody := strings.NewReader(fmt.Sprintf(`{
    "uid": "%v",
    "isEnabled": true,
    "annotationsEnabled": true
  }`, dashboardId))

	PublicDashboardEndpoint := fmt.Sprintf(`/api/dashboards/uid/%v/public-dashboards`, dashboardId)

	resBody, err := SendGrafanaRequest("POST", PublicDashboardEndpoint, PublicDashboardBody)
	if err != nil {
		return "", err
	}

	return resBody["accessToken"].(string), nil
}
