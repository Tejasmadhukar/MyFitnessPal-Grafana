package grafana

import (
	"fmt"
	"testing"
)

func TestCreateDashboard(t *testing.T) {
	fileId := "urdZlJnjRw2OQJeXhQgKf"

	snapShotBody := fmt.Sprintf(`{
    "dashboard": %v,
    "expires": 0,
    "external": false,
    "name": "%v"
  }`, GetDashboard(fileId), fileId)

	t.Log(snapShotBody)
}
