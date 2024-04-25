package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pump struct {
	fuel      int
	line      []*car
	minTime   int
	maxTime   int
	mu        sync.Mutex
	maxInLine int
}

func (carFactory *carFactory) getId() int {
	carFactory.carId += 1
	return carFactory.carId
}

func (pump *Pump) canGetInLine(car *car) bool {
	return pump.fuel == car.fuel && len(pump.line) <= pump.maxInLine
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
			fmt.Printf("car %d fuel: %d left \n", pump.line[0].id, pump.line[0].fuel)
			pump.mu.Lock()
			pump.line = pump.line[1:]
			pump.mu.Unlock()
		case refueling:
			pump.serve()
			pump.line[0].state = lookingForCashier
			fmt.Printf("car %d fuel: %d refilled  \n", pump.line[0].id, pump.line[0].fuel)
		case lookingForCashier:
			gasStation.GetBestCashier().AddToLine(pump.line[0])
		default:
			//do nothing
		}
	}
}
func (pump *Pump) addCar(car *car) {
	pump.mu.Lock()
	pump.line = append(pump.line, car)
	pump.mu.Unlock()
}
func CreatePump(fuel int, maxCarsInLine int, minTime int, maxTime int) *Pump {
	p := new(Pump)
	p.fuel = fuel
	p.line = make([]*car, 0)
	p.maxTime = maxTime
	p.minTime = minTime
	p.maxInLine = maxCarsInLine
	return p
}
