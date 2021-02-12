package models

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
