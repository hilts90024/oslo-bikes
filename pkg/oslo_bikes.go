package pkg

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"oslo-bikes/pkg/models"

	"github.com/caarlos0/env/v6"
)

const (
	SystemInformationPath  = "system_information.json"
	StationInformationPath = "station_information.json"
	StationStatusPath      = "station_status.json"
)

type OsloBikesConfig struct {
	BaseURL          string `env:"BASE_URL,required"`
	ClientIdentifier string `env:"CLIENT_IDENTIFIER,required"`
}

func NewConfig() (*OsloBikesConfig, error) {
	config := OsloBikesConfig{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

type OsloBikesClient struct {
	ctx       context.Context
	apiClient *APIClient
}

func NewOsloBikesClient(ctx context.Context, http *http.Client, cfg *OsloBikesConfig) (*OsloBikesClient, error) {
	apiClient := NewAPIClient(ctx, http, cfg.BaseURL, cfg.ClientIdentifier)

	return &OsloBikesClient{
		ctx:       ctx,
		apiClient: apiClient,
	}, nil
}

func (client *OsloBikesClient) GetAvailableBikesAndDocks() (*models.AvailableBikesDocksResponse, error) {
	stationInfoList, err := client.GetStationInformation()
	if err != nil {
		return nil, err
	}

	stationStatusList, err := client.GetStationStatus()
	if err != nil {
		return nil, err
	}
	stationStatusMap := stationStatusList.Data.AsMap()

	var availableBikesDocksList []*models.AvailableBikesDocks
	for _, stationInfo := range stationInfoList.Data.Stations {
		if stationStatus, ok := stationStatusMap[stationInfo.StationID]; ok {
			availableBikesDocks := &models.AvailableBikesDocks{
				StationID:      stationInfo.StationID,
				StationName:    stationInfo.Name,
				AvailableBikes: stationStatus.NumDocksAvailable,
				AvailableDocks: stationStatus.NumBikesAvailable,
			}
			availableBikesDocksList = append(availableBikesDocksList, availableBikesDocks)
		} else {
			return nil, fmt.Errorf("No station status for ID: %s\n", stationInfo.StationID)
		}
	}

	return &models.AvailableBikesDocksResponse{
		LastUpdated: stationStatusList.LastUpdated,
		Data: &models.AvailableBikesDocksList{
			Stations: availableBikesDocksList,
		},
	}, nil
}

func (c *OsloBikesClient) GetSystemInformation() (*models.SystemInformationResponse, error) {
	body, err := c.apiClient.Get(SystemInformationPath)
	if err != nil {
		return nil, err
	}
	var response = new(models.SystemInformationResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *OsloBikesClient) GetStationInformation() (*models.StationInformationListResponse, error) {
	body, err := c.apiClient.Get(StationInformationPath)
	if err != nil {
		return nil, err
	}
	var response = new(models.StationInformationListResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (c *OsloBikesClient) GetStationStatus() (*models.StationStatusListResponse, error) {
	body, err := c.apiClient.Get("station_status.json")
	if err != nil {
		return nil, err
	}
	var response = new(models.StationStatusListResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response, nil
}
