package main

import (
	"container/list"
	"fmt"
	"math/rand"
)

type car struct {
	id   int
	fuel int
}
type pump struct {
	fuel          int
	line          *list.List
	maxCarsInLine int
}
type carFactory struct {
	carId int
}

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	return car
}
func (carFactory *carFactory) getId() int {
	carFactory.carId += 1
	return carFactory.carId
}
func (pump *pump) canGetInLine(car *car) bool {
	return pump.fuel == car.fuel && pump.line.Len() <= pump.maxCarsInLine
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
	var gaspump pump
	gaspump.fuel = gas
	var testcar = factory.getCar()
	if gaspump.canGetInLine(&testcar) {
		gaspump.line.PushBack(&testcar)
	}

	fmt.Println("Hello, World!")

}
