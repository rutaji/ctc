package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(3)
	fmt.Println("start")
	var factory carFactory
	var gasStation gasStation

	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	var config ConfigManager
	bytes, er := io.ReadAll(jsonFile)
	if er != nil {
		fmt.Println(er)
	}
	json.Unmarshal(bytes, &config)
	var numberOfCars = config.numberOfCars
	gasStation = config.gasStation

	/*gasStation.pumps = append(gasStation.pumps, CreatePump(diesel, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(LPG, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(electric, 10, 10, 40))
	gasStation.cashiers = append(gasStation.cashiers, CreateCashier(5, 10, 20))*/

	gasStation.Open()

	for i := 0; i < numberOfCars; i++ {
		var car = factory.getCar()
		for {
			var bestPump = gasStation.GetBestPump(car.fuel)
			if bestPump == nil {
				fmt.Printf("no pump for %s ,leaving \n", GetFuelName(car.fuel))
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
