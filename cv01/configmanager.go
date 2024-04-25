package main

type ConfigManager struct {
	GasStation   gasStation `json:"gas_station"`
	NumberOfCars int        `json:"number_of_cars"`
}
