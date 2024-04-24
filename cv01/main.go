package main

import (
	"fmt"
	"math/rand"
	"time"
)

var ch chan int

type car struct {
	id    int
	fuel  int
	state int
}
type Pump struct {
	fuel    int
	line    []*car
	minTime int
	maxTime int
}
type carFactory struct {
	carId int
}
type gasStation struct {
	pumps    []*Pump
	cashiers []*Cashier
}
type Cashier struct {
	line    []*car
	minTime int
	maxTime int
}

func (cashier *Cashier) mainLoop() {
	for {
		if len(cashier.line) != 0 {
			var randTime = rand.Intn(cashier.maxTime-cashier.minTime) + cashier.minTime
			time.Sleep(time.Duration(randTime) * time.Millisecond)
			fmt.Printf("car %d paid the price \n", cashier.line[0].id)
			cashier.line[0].state = paid
			cashier.line = cashier.line[1:]
		}
	}

}

func (carFactory *carFactory) getCar() car {
	var car car
	car.id = carFactory.getId()
	car.fuel = rand.Intn(numberOfFuel)
	car.state = refueling
	return car
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
			fmt.Printf("car %d is full \n", pump.line[0].id)
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
func CreateCashier(maxCarsInLine int, minTime int, maxTime int) *Cashier {
	c := new(Cashier)
	c.line = make([]*car, 0, maxCarsInLine)
	c.minTime = minTime
	c.maxTime = maxTime
	return c

}
func (gasStation *gasStation) GetBestPump(fuel int) *Pump {
	var bestPump *Pump
	for _, pump := range gasStation.pumps {
		if pump.fuel != fuel {
			continue
		}
		if bestPump == nil || len(pump.line) < len(bestPump.line) {
			bestPump = pump
		}
	}
	return bestPump
}
func (gasStation *gasStation) Open() {
	for _, pump := range gasStation.pumps {
		go pump.mainLoop(gasStation)
	}
	for _, cashier := range gasStation.cashiers {
		go cashier.mainLoop()
	}
}
func (gasStation *gasStation) EveryOneIsGone() bool {
	for _, pump := range gasStation.pumps {
		if len(pump.line) != 0 {
			return false
		}
	}
	return true
}

const (
	gas = iota
	diesel
	LPG
	electric

	numberOfFuel
)
const (
	refueling = iota
	lookingForCashier
	waitingAtCashier
	paid
)

func main() {
	var factory carFactory
	var gasStation gasStation
	gasStation.pumps = append(gasStation.pumps, CreatePump(gas, 10, 50, 70))
	gasStation.pumps = append(gasStation.pumps, CreatePump(diesel, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(LPG, 10, 10, 40))
	gasStation.pumps = append(gasStation.pumps, CreatePump(electric, 10, 10, 40))
	gasStation.cashiers = append(gasStation.cashiers, CreateCashier(5, 10, 20))
	gasStation.Open()
	var numberOfCars = 6
	for i := 0; i < numberOfCars; i++ {
		var car = factory.getCar()
		car.fuel = gas
		for {
			var bestPump = gasStation.GetBestPump(car.fuel)
			if bestPump == nil {
				fmt.Printf("no pump for %s ,leaving \n", car.fuel)
				break
			}
			if bestPump.canGetInLine(&car) {
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
	fmt.Println("END")

}
