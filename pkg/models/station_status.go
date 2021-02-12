package models

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
