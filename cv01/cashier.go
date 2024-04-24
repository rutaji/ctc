package main

import (
	"fmt"
	"math/rand"
	"time"
)

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

func CreateCashier(maxCarsInLine int, minTime int, maxTime int) *Cashier {
	c := new(Cashier)
	c.line = make([]*car, 0, maxCarsInLine)
	c.minTime = minTime
	c.maxTime = maxTime
	return c

}
