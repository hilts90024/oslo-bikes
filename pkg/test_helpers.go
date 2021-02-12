package pkg

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const stationInformationTestData = `
{
	"last_updated": 1553592653,
	"data": {
		"stations": [
		{
			"station_id":"175",
			"name":"Skøyen Stasjon",
			"address":"Skøyen Stasjon",
			"lat":59.9226729,
			"lon":10.6788129,
			"capacity":20
		},
		{
			"station_id":"47",
			"name":"7 Juni Plassen",
			"address":"7 Juni Plassen",
			"lat":59.9150596,
			"lon":10.7312715,
			"capacity":15
		},
		{
			"station_id":"10",
			"name":"Sotahjørnet",
			"address":"Sotahjørnet",
			"lat":59.9099822,
			"lon":10.7914482,
			"capacity":20
		}
		]
	}
}
`

const stationStatusTestData = `
{
	"last_updated": 1540219230,
	"data": {
		"stations": [
		{
			"is_installed": 1,
			"is_renting": 1,
			"num_bikes_available": 7,
			"num_docks_available": 5,
			"last_reported": 1540219230,
			"is_returning": 1,
			"station_id": "175"
		},
		{
			"is_installed": 1,
			"is_renting": 1,
			"num_bikes_available": 4,
			"num_docks_available": 8,
			"last_reported": 1540219230,
			"is_returning": 1,
			"station_id": "47"
		},
		{
			"is_installed": 1,
			"is_renting": 1,
			"num_bikes_available": 4,
			"num_docks_available": 9,
			"last_reported": 1540219230,
			"is_returning": 1,
			"station_id": "10"
		}
		]
	}
}
`

const systemInformationTestData = `
{
	"last_updated": 1553592653,
	"ttl": 10,
	"data": {
		"system_id": "oslobysykkel",
		"language": "nb",
		"name": "Oslo Bysykkel",
		"operator": "UIP Oslo Bysykkel AS",
		"timezone": "Europe/Oslo",
		"phone_number": "+4791589700",
		"email": "post@oslobysykkel.no"
	}
}
`

type RoundTripFunc func(req *http.Request) *http.Response

func (f RoundTripFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req), nil
}

func NewTestClient(fn RoundTripFunc) *http.Client {
	return &http.Client{
		Transport: RoundTripFunc(fn),
	}
}

func configureHttpTestClient() *http.Client {
	client := NewTestClient(func(req *http.Request) *http.Response {
		var respString string
		url := req.URL
		switch url.Path {
		case fmt.Sprintf("test.no/%s", SystemInformationPath):
			respString = systemInformationTestData
		case fmt.Sprintf("test.no/%s", StationInformationPath):
			respString = stationInformationTestData
		case fmt.Sprintf("test.no/%s", StationStatusPath):
			respString = stationStatusTestData
		default:
			return &http.Response{
				StatusCode: 404,
				Body:       ioutil.NopCloser(bytes.NewBufferString(`Not found`)),
				Header:     make(http.Header),
			}
		}

		return &http.Response{
			StatusCode: 200,
			Body:       ioutil.NopCloser(bytes.NewBufferString(respString)),
			Header:     make(http.Header),
		}
	})
	return client
}
