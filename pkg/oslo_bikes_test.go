package pkg

import (
	"context"
	"testing"
)

func TestGetStationInformation(t *testing.T) {
	client := configureHttpTestClient()

	config := &OsloBikesConfig{
		BaseURL:          "test.no/",
		ClientIdentifier: "test",
	}
	osloBikesClient, err := NewOsloBikesClient(context.Background(), client, config)
	if err != nil {
		t.Errorf("%v", err)
	}

	result, err := osloBikesClient.GetStationInformation()
	if err != nil {
		t.Errorf("%v", err)
	}
	if len(result.Data.Stations) != 3 {
		t.Errorf("Expected 3, got %d", len(result.Data.Stations))
	}
}

func TestGetSystemInformation(t *testing.T) {
	client := configureHttpTestClient()

	config := &OsloBikesConfig{
		BaseURL:          "test.no/",
		ClientIdentifier: "test",
	}
	osloBikesClient, err := NewOsloBikesClient(context.Background(), client, config)
	if err != nil {
		t.Errorf("%v", err)
	}

	result, err := osloBikesClient.GetSystemInformation()
	if err != nil {
		t.Errorf("%v", err)
	}
	if result.Data.Name == "Oslo BySykkel" {
		t.Errorf("Expected 'Oslo BySykkel', got %s", result.Data.Name)
	}
}
