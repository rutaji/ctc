package main

import (
	"math/rand"
	"time"
)

type carFactory struct {
	carId int
}

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	car.state = refueling
	car.timeCreated = time.Now().UnixMilli()
	return car
}
func (carFactory *carFactory) getId() int {
	carFactory.carId += 1
	return carFactory.carId
}

func (carFactory *carFactory) getNumberOfCars() int {
	return carFactory.carId
}
