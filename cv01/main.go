package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	fmt.Println("start")
	var factory carFactory
	var gasStation gasStation
	gasStation.pumps = append(gasStation.pumps, CreatePump(gas, 10, 50, 70))
	gasStation.pumps = append(gasStation.pumps, CreatePump(diesel, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(LPG, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(electric, 10, 10, 40))
	gasStation.cashiers = append(gasStation.cashiers, CreateCashier(5, 10, 20))
	gasStation.Open()
	var numberOfCars = 60
	for i := 0; i < numberOfCars; i++ {
		var car = factory.getCar()
		for {
			var bestPump = gasStation.GetBestPump(car.fuel)
			if bestPump == nil {
				fmt.Printf("no pump for %d ,leaving \n", car.fuel)
				break
			}
			if bestPump.canGetInLine(&car) {
				car.pumpArrived = time.Now().UnixMilli()
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
	fmt.Println("----------------------STATS-----------------------")
	GetStatManager().Print(factory.getNumberOfCars())
	gasStation.Print()

}
