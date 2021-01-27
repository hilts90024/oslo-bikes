package main

import (
	"context"
	"fmt"
	"log"
	"oslo-bikes/pkg"
)

func Main() error {
	config, err := pkg.NewConfig()
	if err != nil {
		return err
	}

	client, err := pkg.NewOsloBikesClient(context.Background(), config)
	if err != nil {
		return err
	}

	systemInfo, err := client.GetSystemInformation()
	if err != nil {
		return err
	}
	stationInfoList, err := client.GetStationInformation()
	if err != nil {
		return err
	}

	stationStatusList, err := client.GetStationStatus()
	if err != nil {
		return err
	}
	stationStatusMap := stationStatusList.AsMap()

	fmt.Printf("System: %s\n", systemInfo.Name)
	for _, stationInfo := range stationInfoList.Stations {
		if stationStatus, ok := stationStatusMap[stationInfo.StationID]; ok {
			fmt.Printf("Station: %s, available docks %d, available bikes %d\n",
				stationInfo.Name,
				stationStatus.NumDocksAvailable,
				stationStatus.NumBikesAvailable)
		} else {
			return fmt.Errorf("No station info for ID: %s\n", stationStatus.StationID)
		}
	}
	fmt.Printf("No more stations in system: %s\n", systemInfo.Name)
	return nil

}
func main() {
	if err := Main(); err != nil {
		log.Fatalf("%v", err)
	}
}
