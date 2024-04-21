package main

import (
	"fmt"
	"math/rand"
	"time"
)

type car struct {
	id         int
	fuel       int
	isTankFull bool
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
	pumps []*pump

}
type 

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	car.isTankFull = false
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
func (pump *pump) mainLoop() {
	for {
		if !pump.line[0].isTankFull {
			pump.serve()
			pump.line[0].isTankFull = true
			fmt.Printf("car %d is full", pump.line[0].id)
		}
		//waiting for payment
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
