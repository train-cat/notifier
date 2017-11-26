package api

import (
	"github.com/train-cat/client-train-go"
	"github.com/train-cat/client-train-go/filters"
)

// GetAlerts return all alerts for one stop
func GetAlerts(code string, station int) ([]traincat.Alert, error) {
	f := &filters.Alert{
		CodeTrain: &code,
		StationID: &station,
	}

	return traincat.CGetAllAlerts(f)
}
