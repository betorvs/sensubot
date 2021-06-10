package opsgenieclient

import (
	"context"
	"strings"
	"time"

	"github.com/betorvs/sensubot/appcontext"
	"github.com/betorvs/sensubot/config"
	"github.com/opsgenie/opsgenie-go-sdk-v2/alert"
	"github.com/opsgenie/opsgenie-go-sdk-v2/client"
)

// OpsgenieRepository struct
type OpsgenieRepository struct {
	Alert *alert.Client
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
		return getResult, nil
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
		return listResult, nil
	}
	return listResult, nil
}

func New() appcontext.Component {
	alertClient, _ := alert.NewClient(&client.Config{
		ApiKey:         config.Values.OpsgenieAPIKey,
		OpsGenieAPIURL: switchOpsgenieRegion(),
	})
	return &OpsgenieRepository{Alert: alertClient}
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
