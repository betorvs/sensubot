package usecase

import (
	"fmt"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
)

// requestOpsgenie func receives a map[string]string command, role and display Name and process it using Opsgenie api
func requestOpsgenie(command map[string]string, role, displayName string) string {
	if len(config.Values.OpsgenieAPIKey) == 0 {
		return "SensuBot Not Configured to access any Opsgenie"
	}
	var s string
	logLocal := config.GetLogger()
	switch command["verb"] {
	case "list":
		switch command["resource"] {
		case "alerts":
			query := "status=open"
			results, err := listAlerts(query)
			if err != nil {
				logLocal.Error(err)
			}
			var r string
			for _, v := range results {
				r += fmt.Sprintf("ID: %s, Alert Name: %s, Priority: %v \n", v.TinyID, v.Message, v.Priority)
			}
			s += fmt.Sprintf("Number of Alerts found: %v \n", len(results))
			s += r

			// end of command resource
		}
	case "get":
		switch command["resource"] {
		case "alerts":
			result, err := getAlerts(command["name"])
			if err != nil {
				logLocal.Error(err)
			}
			s += fmt.Sprintf("ID: %s, Alert Name: %s, Priority: %v \n", result.TinyId, result.Message, result.Priority)
			s += fmt.Sprintf("Description:\n %s \n", result.Description)

			// end of command resource
		}
		// end of command verb
	}
	return s
}

func listAlerts(query string) ([]alert.Alert, error) {
	repo := domain.GetOpsgenieRepository()
	results, err := repo.ListAlert(query)
	if err != nil {
		return results.Alerts, err
	}
	return results.Alerts, nil
}

func getAlerts(tinyID string) (*alert.GetAlertResult, error) {
	repo := domain.GetOpsgenieRepository()
	result, err := repo.GetAlert(tinyID)
	if err != nil {
		return result, err
	}
	return result, nil
}