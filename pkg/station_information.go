package pkg

// {
// 	"last_updated": 1553592653,
// 	"data": {
// 	  "stations": [
// 		{
// 		  "station_id":"627",
// 		  "name":"Skøyen Stasjon",
// 		  "address":"Skøyen Stasjon",
// 		  "lat":59.9226729,
// 		  "lon":10.6788129,
// 		  "capacity":20
// 		},
// 		{
// 		  "station_id":"623",
// 		  "name":"7 Juni Plassen",
// 		  "address":"7 Juni Plassen",
// 		  "lat":59.9150596,
// 		  "lon":10.7312715,
// 		  "capacity":15
// 		},
// 		{
// 		  "station_id":"610",
// 		  "name":"Sotahjørnet",
// 		  "address":"Sotahjørnet",
// 		  "lat":59.9099822,
// 		  "lon":10.7914482,
// 		  "capacity":20
// 		}
// 	  ]
// 	}
// }

type StationInformationListResponse struct {
	LastUpdated uint32 `json:"last_updated"`
	Data        *StationInformationList
}

type StationInformationList struct {
	Stations []*StationInformation
}

type StationInformation struct {
	StationID string `json:"station_id"`
	Name      string
	Address   string
	Lat       float32
	Lon       float32
	Capacity  int
}
