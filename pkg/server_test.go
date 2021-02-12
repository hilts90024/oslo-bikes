package pkg

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http/httptest"
	"oslo-bikes/pkg/models"
	"testing"
)

func TestGetAvailableBikesAndDocks(t *testing.T) {
	// setup
	client := configureHttpTestClient()

	config := &OsloBikesConfig{
		BaseURL:          "test.no/",
		ClientIdentifier: "test",
	}
	osloBikesClient, err := NewOsloBikesClient(context.Background(), client, config)
	if err != nil {
		t.Errorf("%v", err)
	}

	server := NewServer(context.Background(), 8080, osloBikesClient)

	req := httptest.NewRequest("GET", "http://test.no/", nil)
	w := httptest.NewRecorder()

	// test
	server.GetAvailableBikesAndDocks(w, req)

	resp := w.Result()
	body, _ := ioutil.ReadAll(resp.Body)

	if resp.StatusCode != 200 {
		t.Errorf("Expected response 200, got %d", resp.StatusCode)
	}

	var response = new(models.AvailableBikesDocksResponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		t.Errorf("Failed parsing response, %v", err)
	}

	if len(response.Data.Stations) != 3 {
		t.Errorf("Expected 3, got %d", len(response.Data.Stations))
	}
}
