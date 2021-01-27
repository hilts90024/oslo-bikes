package pkg

import (
	"context"
	"encoding/json"

	"github.com/caarlos0/env/v6"
)

type Config struct {
	BaseURL          string `env:"BASE_URL,required"`
	ClientIdentifier string `env:"CLIENT_IDENTIFIER,required"`
}

func NewConfig() (*Config, error) {
	config := Config{}
	if err := env.Parse(&config); err != nil {
		return nil, err
	}
	return &config, nil
}

type OsloBikesClient struct {
	ctx       context.Context
	apiClient *APIClient
}

func NewOsloBikesClient(ctx context.Context, cfg *Config) (*OsloBikesClient, error) {
	apiClient := NewAPIClient(ctx, cfg.BaseURL, cfg.ClientIdentifier)

	return &OsloBikesClient{
		ctx:       ctx,
		apiClient: apiClient,
	}, nil
}

func (c *OsloBikesClient) GetSystemInformation() (*SystemInformation, error) {
	body, err := c.apiClient.Get("system_information.json")
	if err != nil {
		return nil, err
	}
	var response = new(SystemInformationResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *OsloBikesClient) GetStationInformation() (*StationInformationList, error) {
	body, err := c.apiClient.Get("station_information.json")
	if err != nil {
		return nil, err
	}
	var response = new(StationInformationListResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}

func (c *OsloBikesClient) GetStationStatus() (*StationStatusList, error) {
	body, err := c.apiClient.Get("station_status.json")
	if err != nil {
		return nil, err
	}
	var response = new(StationStatusListResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	return response.Data, nil
}
