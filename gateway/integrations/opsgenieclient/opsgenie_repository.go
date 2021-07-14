package opsgenieclient

import (
	"context"
	"strings"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
	"github.com/opsgenie/opsgenie-go-sdk-v2/incident"
	"github.com/opsgenie/opsgenie-go-sdk-v2/schedule"
)

// OpsgenieRepository struct
type OpsgenieRepository struct {
	Alert    *alert.Client
	Incident *incident.Client
	Schedule *schedule.Client
}

// switchOpsgenieRegion func
func switchOpsgenieRegion() client.ApiUrl {
	var region client.ApiUrl
	apiRegionLowCase := strings.ToLower(config.Values.OpsgenieRegion)
	switch apiRegionLowCase {
	case "eu":
		region = client.API_URL_EU
	case "us":
		region = client.API_URL
	default:
		region = client.API_URL
	}
	return region
}

// GetAlert by tinyID
func (repo OpsgenieRepository) GetAlert(tinyID string) (*alert.GetAlertResult, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	getResult, err := repo.Alert.Get(ctx, &alert.GetAlertRequest{
		IdentifierType:  alert.TINYID,
		IdentifierValue: tinyID,
	})
	if err != nil {
		return getResult, err
	}
	return getResult, nil
}

// ListAlert func
func (repo OpsgenieRepository) ListAlert(query string) (*alert.ListAlertResult, error) {
	logLocal := config.GetLogger()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logLocal.Debug("starting list alert ", query)
	listResult, err := repo.Alert.List(ctx, &alert.ListAlertRequest{
		Query: query,
	})
	if err != nil {
		return listResult, err
	}
	return listResult, nil
}

// ListIncident func
func (repo OpsgenieRepository) ListIncident(query string) (*incident.ListResult, error) {
	logLocal := config.GetLogger()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logLocal.Debug("starting list incidents ", query)
	listResult, err := repo.Incident.List(ctx, &incident.ListRequest{
		Query: query,
	})
	if err != nil {
		return listResult, err
	}
	return listResult, nil
}

// CreateIncident func
// func (repo OpsgenieRepository) CreateIncident() error {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()
// 	createResult, err := repo.Incident.
// 	// Create(ctx, &incident.CreateRequest{})
// }

// GetOnCall func
func (repo OpsgenieRepository) GetOnCall(team string) (*schedule.GetOnCallsResult, error) {
	logLocal := config.GetLogger()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logLocal.Debug("geting on calls ", team)
	date := time.Now()
	flat := true
	scheduleResult, err := repo.Schedule.GetOnCalls(ctx, &schedule.GetOnCallsRequest{
		Flat:                   &flat,
		Date:                   &date,
		ScheduleIdentifierType: schedule.Name,
		ScheduleIdentifier:     team,
	})
	if err != nil {
		return scheduleResult, err
	}
	return scheduleResult, nil
}

// ListSchedules func
func (repo OpsgenieRepository) ListSchedules() (*schedule.ListResult, error) {
	logLocal := config.GetLogger()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	logLocal.Debug("ListSchedules starts ")
	expand := true
	scheduleResult, err := repo.Schedule.List(ctx, &schedule.ListRequest{
		Expand: &expand,
	})
	if err != nil {
		return scheduleResult, err
	}
	return scheduleResult, nil

}

func New() appcontext.Component {
	alertClient, _ := alert.NewClient(&client.Config{
		ApiKey:         config.Values.OpsgenieAPIKey,
		OpsGenieAPIURL: switchOpsgenieRegion(),
	})
	incidentClient, _ := incident.NewClient(&client.Config{
		ApiKey:         config.Values.OpsgenieAPIKey,
		OpsGenieAPIURL: switchOpsgenieRegion(),
	})
	scheduleClient, _ := schedule.NewClient(&client.Config{
		ApiKey:         config.Values.OpsgenieAPIKey,
		OpsGenieAPIURL: switchOpsgenieRegion(),
	})
	return &OpsgenieRepository{
		Alert:    alertClient,
		Incident: incidentClient,
		Schedule: scheduleClient,
	}
}

func init() {
	if config.Values.TestRun {
		return
	}

	if config.Values.OpsgenieAPIKey != "Absent" {
		appcontext.Current.Add(appcontext.OpsgenieRepository, New)
		logLocal := config.GetLogger()
		logLocal.Info("Connected to Opsgenie")
	}

}
