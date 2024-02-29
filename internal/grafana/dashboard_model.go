package grafana

import "fmt"

func GetDashboard(fileuid string) string {
	dashboardModel := fmt.Sprintf(`{
    "annotations": {
      "list": [
        {
          "builtIn": 1,
          "datasource": {
            "type": "grafana",
            "uid": "-- Grafana --"
          },
          "enable": true,
          "hide": true,
          "iconColor": "rgba(0, 211, 255, 1)",
          "name": "Annotations & Alerts",
          "type": "dashboard"
        }
      ]
    },
    "editable": true,
    "fiscalYearStartMonth": 0,
    "graphTooltip": 0,
    "id": null,
    "links": [],
    "liveNow": false,
    "panels": [
      {
        "datasource": {
          "type": "marcusolsson-csv-datasource",
          "uid": "%v"
        },
        "description": "It shows, how many calories you have consumed over time",
        "fieldConfig": {
          "defaults": {
            "color": {
              "mode": "palette-classic"
            },
            "custom": {
              "axisBorderShow": false,
              "axisCenteredZero": false,
              "axisColorMode": "text",
              "axisLabel": "",
              "axisPlacement": "auto",
              "barAlignment": 0,
              "drawStyle": "line",
              "fillOpacity": 0,
              "gradientMode": "none",
              "hideFrom": {
                "legend": false,
                "tooltip": false,
                "viz": false
              },
              "insertNulls": false,
              "lineInterpolation": "linear",
              "lineWidth": 1,
              "pointSize": 5,
              "scaleDistribution": {
                "type": "linear"
              },
              "showPoints": "auto",
              "spanNulls": false,
              "stacking": {
                "group": "A",
                "mode": "none"
              },
              "thresholdsStyle": {
                "mode": "off"
              }
            },
            "mappings": [],
            "thresholds": {
              "mode": "absolute",
              "steps": [
                {
                  "color": "green",
                  "value": null
                },
                {
                  "color": "red",
                  "value": 80
                }
              ]
            },
            "unitScale": true
          },
          "overrides": []
        },
        "gridPos": {
          "h": 8,
          "w": 12,
          "x": 2,
          "y": 0
        },
        "id": 1,
        "options": {
          "legend": {
            "calcs": [],
            "displayMode": "list",
            "placement": "bottom",
            "showLegend": true
          },
          "tooltip": {
            "mode": "single",
            "sort": "none"
          }
        },
        "pluginVersion": "10.3.3",
        "targets": [
          {
            "datasource": {
              "type": "marcusolsson-csv-datasource",
              "uid": "%v"
            },
            "decimalSeparator": ".",
            "delimiter": ",",
            "header": true,
            "ignoreUnknown": true,
            "refId": "A",
            "schema": [
              {
                "name": "Date",
                "type": "time"
              },
              {
                "name": "Calories",
                "type": "number"
              }
            ],
            "skipRows": 0
          }
        ],
        "title": "Calories vs Time",
        "type": "timeseries"
      },
      {
        "datasource": {
          "type": "marcusolsson-csv-datasource",
          "uid": "%v"
        },
        "description": "Average amount of calories consumed over time.",
        "fieldConfig": {
          "defaults": {
            "mappings": [],
            "thresholds": {
              "mode": "absolute",
              "steps": [
                {
                  "color": "green",
                  "value": null
                },
                {
                  "color": "red",
                  "value": 80
                }
              ]
            },
            "unit": "short",
            "unitScale": true
          },
          "overrides": []
        },
        "gridPos": {
          "h": 7,
          "w": 7,
          "x": 14,
          "y": 0
        },
        "id": 2,
        "options": {
          "colorMode": "value",
          "graphMode": "area",
          "justifyMode": "auto",
          "orientation": "auto",
          "reduceOptions": {
            "calcs": [
              "mean"
            ],
            "fields": "",
            "values": false
          },
          "showPercentChange": false,
          "textMode": "auto",
          "wideLayout": true
        },
        "pluginVersion": "10.3.3",
        "targets": [
          {
            "datasource": {
              "type": "marcusolsson-csv-datasource",
              "uid": "%v"
            },
            "decimalSeparator": ".",
            "delimiter": ",",
            "header": true,
            "ignoreUnknown": true,
            "refId": "A",
            "schema": [
              {
                "name": "Date",
                "type": "time"
              },
              {
                "name": "Calories",
                "type": "number"
              },
              {
                "name": "Protien (g)",
                "type": "number"
              }
            ],
            "skipRows": 0
          }
        ],
        "title": "Average Calories",
        "type": "stat"
      }
    ],
    "refresh": "",
    "schemaVersion": 39,
    "tags": [],
    "templating": {
      "list": []
    },
    "timepicker": {},
    "timezone": "",
    "title": "%v",
    "uid": null,
    "version": 2,
    "weekStart": ""}`, fileuid, fileuid, fileuid, fileuid, fileuid)
	return dashboardModel
}
