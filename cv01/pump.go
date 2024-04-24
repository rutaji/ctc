package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Pump struct {
	fuel    int
	line    []*car
	minTime int
	maxTime int
}

func (carFactory *carFactory) getId() int {
	carFactory.carId += 1
	return carFactory.carId
}

func (pump *Pump) canGetInLine(car *car) bool {
	return pump.fuel == car.fuel && len(pump.line) <= cap(pump.line)
}
func (pump *Pump) serve() {
	var randTime = rand.Intn(pump.maxTime-pump.minTime) + pump.minTime
	time.Sleep(time.Duration(randTime) * time.Millisecond)
}
func (pump *Pump) mainLoop(gasStation *gasStation) {
	for {
		if len(pump.line) == 0 {
			continue
		}
		switch pump.line[0].state {
		case paid:
			fmt.Printf("car %d left \n", pump.line[0].id)
			pump.line = pump.line[1:]
		case refueling:
			pump.serve()
			pump.line[0].state = lookingForCashier
			fmt.Printf("car %d refilled \n", pump.line[0].id)
		case lookingForCashier:
			for _, element := range gasStation.cashiers {
				if len(element.line) < cap(element.line) { //todo get best cashier
					pump.line[0].state = waitingAtCashier
					element.line = append(element.line, pump.line[0])
					break
				}
			}
		default:
			//do nothing
		}
	}
}
func (pump *Pump) addCar(car *car) {
	pump.line = append(pump.line, car)
}
func CreatePump(fuel int, maxCarsInLine int, minTime int, maxTime int) *Pump {
	p := new(Pump)
	p.fuel = fuel
	p.line = make([]*car, 0, maxCarsInLine)
	p.maxTime = maxTime
	p.minTime = minTime
	return p
}
