package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch chan int

type car struct {
	id         int
	fuel       int
	isTankFull bool
	isPayed    bool
}
type pump struct {
	fuel          int
	line          []*car
	maxCarsInLine int
	minTime       int
	maxTime       int
}
type carFactory struct {
	carId int
}
type gasStation struct {
	pumps    []*pump
	cashiers []*cashier
}
type cashier struct {
	payingCar *car
	minTime   int
	maxTime   int
}

func (cashier *cashier) mainLoop() {
	if cashier.payingCar != nil {
		var randTime = rand.Intn(pump.maxTime-pump.minTime) + pump.minTime
		time.Sleep(randTime)
		cashier.payingCar.isPayed = true
		fmt.Printf("car %d paid the price", pump.line[0].id)
		cashier.payingCar = nil
	}

}

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	car.isTankFull = false
	car.isPayed = false
	return car
}
func (carFactory *carFactory) getId() int {
	carFactory.carId += 1
	return carFactory.carId
}
func (pump *pump) canGetInLine(car *car) bool {
	return pump.fuel == car.fuel && len(pump.line) <= pump.maxCarsInLine
}
func (pump *pump) getCarsInLine() int {
	return len(pump.line)
}
func (pump *pump) serve() {
	var randTime = rand.Intn(pump.maxTime-pump.minTime) + pump.minTime
	time.Sleep(time.Duration(randTime) * time.Millisecond)
}
func (pump *pump) mainLoop(gasStation *gasStation) {
	for {
		if pump.line[0].isPayed {

		}
		else{
			if !pump.line[0].isTankFull {
				pump.serve()
				pump.line[0].isTankFull = true
				fmt.Printf("car %d is full", pump.line[0].id)
			}
			//waiting for payment
			for _, element := range gasStation.cashiers {
				if element.payingCar == nil {
					element.payingCar = pump.line[0]
					break
				}
			}
		}
	}
}
func (pump *pump) addCar() {}
func CreatePump(fuel int, maxCarsInLine int, minTime int, maxTime int) *pump {
	p := new(pump)
	p.fuel = fuel
	p.maxCarsInLine = maxCarsInLine
	p.line = make([]*car, 0, maxCarsInLine)
	p.maxTime = maxTime
	p.minTime = minTime
	return p
}

const (
	gas = iota
	diesel
	LPG
	electric

	numberOfFuel
)

func main() {
	var factory carFactory
	var gaspump = CreatePump(gas, 10, 50, 70)
	var testcar = factory.getCar()
	testcar.fuel = gas

	fmt.Println("Hello, World!")

}
