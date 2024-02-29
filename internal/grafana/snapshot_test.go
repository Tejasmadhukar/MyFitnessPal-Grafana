package grafana

import (
	"fmt"
	"testing"
)

func TestCreateSnapShot(t *testing.T) {
	fileId := "MZwL4VemntraAvWMBpzjr"

	snapShotBody := fmt.Sprintf(`{
    "dashboard": %v,
    "expires": 0,
    "external": false,
    "name": "%v"
  }`, GetDashboard(fileId), fileId)

	t.Log(snapShotBody)
}
