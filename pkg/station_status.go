package pkg

// {
// 	"last_updated": 1540219230,
// 	"data": {
// 	  "stations": [
// 		{
// 		  "is_installed": 1,
// 		  "is_renting": 1,
// 		  "num_bikes_available": 7,
// 		  "num_docks_available": 5,
// 		  "last_reported": 1540219230,
// 		  "is_returning": 1,
// 		  "station_id": "175"
// 		},
// 		{
// 		  "is_installed": 1,
// 		  "is_renting": 1,
// 		  "num_bikes_available": 4,
// 		  "num_docks_available": 8,
// 		  "last_reported": 1540219230,
// 		  "is_returning": 1,
// 		  "station_id": "47"
// 		},
// 		{
// 		  "is_installed": 1,
// 		  "is_renting": 1,
// 		  "num_bikes_available": 4,
// 		  "num_docks_available": 9,
// 		  "last_reported": 1540219230,
// 		  "is_returning": 1,
// 		  "station_id": "10"
// 		}
// 	  ]
// 	}
// }

type StationStatusListResponse struct {
	LastUpdated uint32 `json:"last_updated"`
	Data        *StationStatusList
}

type StationStatusList struct {
	Stations []*StationStatus
}

func (ssl *StationStatusList) AsMap() map[string]*StationStatus {
	stationStatusMap := make(map[string]*StationStatus)
	for _, stationStatus := range ssl.Stations {
		stationStatusMap[stationStatus.StationID] = stationStatus
	}
	return stationStatusMap
}

type StationStatus struct {
	IsInstalled       int    `json:"is_installed"`
	IsRenting         int    `json:"is_renting"`
	NumBikesAvailable int    `json:"num_bikes_available"`
	NumDocksAvailable int    `json:"num_docks_available"`
	LastReported      uint32 `json:"last_reported"`
	IsReturning       int    `json:"is_returning"`
	StationID         string `json:"station_id"`
}
