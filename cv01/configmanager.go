package main

type ConfigManager struct {
	gasStation   gasStation `json:"gas_station"`
	numberOfCars int        `json:"number_of_cars"`
}
