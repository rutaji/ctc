package main

import "math/rand"

type carFactory struct {
	carId int
}

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	car.state = refueling
	return car
}
