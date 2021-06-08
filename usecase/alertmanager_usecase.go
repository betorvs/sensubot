package usecase

import (
	"encoding/json"
	"fmt"

	"github.com/betorvs/sensubot/config"
	"github.com/betorvs/sensubot/domain"
	"github.com/betorvs/sensubot/utils"
	"github.com/prometheus/alertmanager/api/v2/models"
)

// requestAlertManager func receives a map[string]string command, role and display Name and process it using AlertManager api
func requestAlertManager(command map[string]string, role, displayName string) string {
	if len(config.Values.AlertManagerEndpoints) == 0 {
		return "SensuBot Not Configured to access any Alert Manager Endpoint"
	}

	var s string
	logLocal := config.GetLogger()
	switch command["verb"] {
	case "list":
		var b bool
		labels := make(map[string]string)
		if command["labels"] != "" {
			k, v, t := utils.ParseKeyValue(command["labels"], "=")
			b = t
			labels[k] = v
		}
		if !b {
			output := "Sorry. List request should be used with labels"
			return output
		}
		var r string
		switch command["resource"] {
		case "alerts":
			count := 0
			for _, am := range config.Values.AlertManagerEndpoints {
				amURL := createAlertManagerURI(am, "alerts")
				alerts, err := getAlertManagerAlerts(amURL, labels)
				if err == nil {
					for _, a := range alerts {
						if a.Labels["alertname"] != "" {
							if a.Labels["job"] != "" {
								r += fmt.Sprintf("Alert Name: %s, Job: %s \n", a.Labels["alertname"], a.Labels["job"])
							} else {
								r += fmt.Sprintf("Alert Name: %s \n", a.Labels["alertname"])
							}
							count++
						}
					}
				}
			}
			s += fmt.Sprintf("Number of Alerts found: %v \n", count)
			s += r
		case "silences":
			count := 0
			for _, am := range config.Values.AlertManagerEndpoints {
				amURL := createAlertManagerURI(am, "silences")
				silences, err := getAlertManagerSilences(amURL, labels)
				if err == nil {
					for _, s := range silences {
						r += fmt.Sprintf("Silence ID: %s , Created by: %s , State: %s \n", *s.ID, *s.CreatedBy, *s.Status.State)
						count++
					}
				}
			}
			s += fmt.Sprintf("Number of Silences found: %v \n", count)
			s += r
		}

	case "get":
		logLocal.Debug("alertmanager resource get")
		var b bool
		labels := make(map[string]string)
		if command["labels"] != "" {
			k, v, t := utils.ParseKeyValue(command["labels"], "=")
			b = t
			labels[k] = v
		}
		if !b {
			output := "Sorry. List request should be used with labels"
			return output
		}
		var r string
		switch command["resource"] {
		case "alerts":
			count := 0
			for _, am := range config.Values.AlertManagerEndpoints {
				amURL := createAlertManagerURI(am, "alerts")
				alerts, err := getAlertManagerAlerts(amURL, labels)
				if err == nil {
					for _, a := range alerts {
						r += fmt.Sprintf("- Alert Name: %s \n  Labels: \n", a.Labels["alertname"])
						for k, v := range a.Labels {
							if k != "alertname" {
								r += fmt.Sprintf("  - %s: %s \n", k, v)
							}
						}
						r += "  Annotations: \n"
						for k, v := range a.Annotations {
							r += fmt.Sprintf("  - %s: %s \n", k, v)
						}
						count++
					}
				}
			}
			s += fmt.Sprintf("Number of Alerts found: %v \n", count)
			s += r
		case "silences":
			count := 0
			for _, am := range config.Values.AlertManagerEndpoints {
				amURL := createAlertManagerURI(am, "silences")
				silences, err := getAlertManagerSilences(amURL, labels)
				if err == nil {
					for _, s := range silences {
						r += fmt.Sprintf("Silence ID: %s , Created by: %s, State: %s \n  - Comment: %s \n", *s.ID, *s.CreatedBy, *s.Status.State, *s.Comment)
						r += "  - Matchers:\n"
						for _, v := range s.Matchers {
							r += fmt.Sprintf("  - %s : %s \n", *v.Name, *v.Value)
						}
						count++
					}
				}
			}
			s += fmt.Sprintf("Number of Silences found: %v \n", count)
			s += r
		}
	}
	logLocal.Debug(s)
	return s
}

func createAlertManagerURI(am, kind string) string {

	switch kind {
	case "alerts":
		return fmt.Sprintf("%s/api/v2/alerts", am)
	case "silences":
		return fmt.Sprintf("%s/api/v2/silences", am)
	}
	return ""
}

func getAlertManagerAlerts(amURL string, labelsSelector map[string]string) ([]models.GettableAlert, error) {
	repo := domain.GetAlertManagerRepository()
	body, err := repo.GetAlertManager(amURL)
	alerts := []models.GettableAlert{}
	if err != nil {
		return alerts, fmt.Errorf("Failed to get alert manager alerts: %v", err)
	}

	_ = json.Unmarshal(body, &alerts)

	result := filterAlerts(alerts, labelsSelector)

	return result, nil
}

// filter alerts using map[string]string from LabelsSelector
func filterAlerts(alerts []models.GettableAlert, labelsSelector map[string]string) (result []models.GettableAlert) {
	for _, alert := range alerts {
		for key, value := range labelsSelector {
			if alert.Labels[key] == value {
				result = append(result, alert)
			}
		}
	}
	return result
}

func getAlertManagerSilences(amURL string, labelsSelector map[string]string) ([]models.GettableSilence, error) {
	repo := domain.GetAlertManagerRepository()
	body, err := repo.GetAlertManager(amURL)
	silences := []models.GettableSilence{}
	if err != nil {
		return silences, fmt.Errorf("Failed to get alert manager silences: %v", err)
	}

	_ = json.Unmarshal(body, &silences)

	result := filterSilences(silences, labelsSelector)

	return result, nil
}

func filterSilences(silences []models.GettableSilence, labelsSelector map[string]string) (result []models.GettableSilence) {
	for _, silence := range silences {
		for key, value := range labelsSelector {
			for _, m := range silence.Matchers {
				if *m.Name == key && *m.Value == value {
					result = append(result, silence)
				}
			}
		}
	}
	return result
}
