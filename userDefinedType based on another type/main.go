package main

import (
	"fmt"
)

type KathmanduTravels struct {
	altitude int
	location string
}

func (k KathmanduTravels) showPollution(pollutionLevel int) int {
	return k.altitude + pollutionLevel
}

func (k KathmanduTravels) showHills() string {
	return "Hilly Location:" + k.location
}
type LouisvilleTravels KathmanduTravels

func (k LouisvilleTravels) showMallsCount(mallsCount int) int {
	return k.altitude + mallsCount
}

func (k LouisvilleTravels) showRoads() string {
	return "Road Location:" + k.location
}

func main() {
	kathmanduTravels:=KathmanduTravels{
		altitude: 1200,
		location: "Chandragiri",
	}

	fmt.Println(kathmanduTravels.showPollution(1))
	fmt.Println(kathmanduTravels.showHills())

	louisvilleTravels:=LouisvilleTravels{
		altitude: 400,
		location: "Chandragiri",
	}

	fmt.Println(louisvilleTravels.showMallsCount(1))
	fmt.Println(louisvilleTravels.showRoads())
}
