package pkg

import (
	"context"
	"io/ioutil"
	"net/http"
)

type APIClient struct {
	ctx              context.Context
	ClientIdentifier string
	BaseURL          string
	client           *http.Client
}

func NewAPIClient(ctx context.Context, http *http.Client, baseURL string, clientIdentifier string) *APIClient {
	return &APIClient{
		ctx:              ctx,
		BaseURL:          baseURL,
		ClientIdentifier: clientIdentifier,
		client:           http,
	}
}

func (apiClient *APIClient) Get(path string) ([]byte, error) {
	url := apiClient.BaseURL + path
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(apiClient.ctx)
	req.Header.Add("Client-Identifier", apiClient.ClientIdentifier)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")

	resp, err := apiClient.client.Do(req)
	if err != nil {
		return nil, err
	}
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respBody, nil
}
