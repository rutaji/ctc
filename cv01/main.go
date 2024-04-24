package main

import (
	"fmt"
)

func main() {
	var factory carFactory
	var gasStation gasStation
	gasStation.pumps = append(gasStation.pumps, CreatePump(gas, 10, 50, 70))
	gasStation.pumps = append(gasStation.pumps, CreatePump(diesel, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(LPG, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(electric, 10, 10, 40))
	gasStation.cashiers = append(gasStation.cashiers, CreateCashier(5, 10, 20))
	gasStation.Open()
	var numberOfCars = 6
	for i := 0; i < numberOfCars; i++ {
		var car = factory.getCar()
		car.fuel = gas
		for {
			var bestPump = gasStation.GetBestPump(car.fuel)
			if bestPump == nil {
				fmt.Printf("no pump for %s ,leaving \n", car.fuel)
				break
			}
			if bestPump.canGetInLine(&car) {
				bestPump.addCar(&car)
				break
			}

		}

	}
	for {
		if gasStation.EveryOneIsGone() {
			break
		}
	}
	fmt.Println("END")

}
