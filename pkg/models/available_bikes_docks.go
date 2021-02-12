package models

type AvailableBikesDocksResponse struct {
	LastUpdated uint32                   `json:"last_updated"`
	Data        *AvailableBikesDocksList `json:"data"`
}

type AvailableBikesDocksList struct {
	Stations []*AvailableBikesDocks `json:"stations"`
}

type AvailableBikesDocks struct {
	StationID      string `json:"station_id"`
	StationName    string `json:"station_name"`
	AvailableBikes int    `json:"available_bikes"`
	AvailableDocks int    `json:"available_docks"`
}
