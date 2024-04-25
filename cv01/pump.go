package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pump struct {
	fuel      int `json:"fuel"`
	line      []*car
	minTime   int `json:"min_time"`
	maxTime   int `json:"max_time"`
	mu        sync.Mutex
	maxInLine int `json:"max_in_line"`

	totalTimeinLine int64
	totalcars       int
}

func (pump *Pump) Print() {
	fmt.Printf("Pump fuel: %s\n	total time in line: %d s\n	total cars: %d \n	avg. time in line %d ms\n",
		GetFuelName(pump.fuel), pump.totalTimeinLine/1000, pump.totalcars, pump.totalTimeinLine/int64(pump.totalcars))
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
			fmt.Printf("car %d fuel: %s left \n", pump.line[0].id, GetFuelName(pump.line[0].fuel))
			GetStatManager().CarLeft(pump.line[0].timeCreated)
			pump.totalcars++
			pump.mu.Lock()
			pump.line = pump.line[1:]
			pump.mu.Unlock()
		case refueling:
			pump.totalTimeinLine += time.Now().UnixMilli() - pump.line[0].pumpArrived
			pump.serve()
			pump.line[0].state = lookingForCashier
			fmt.Printf("car %d fuel: %s refilled  \n", pump.line[0].id, GetFuelName(pump.line[0].fuel))
			pump.line[0].cashierArrived = time.Now().UnixMilli()
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
func (p *Pump) CreatePump() {
	p.line = make([]*car, 0)
}
