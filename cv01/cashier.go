package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cashier struct {
	line      []*car
	minTime   int
	maxTime   int
	mu        sync.Mutex
	maxInLine int
}

func (cashier *Cashier) mainLoop() {
	for {
		if len(cashier.line) != 0 {
			var randTime = rand.Intn(cashier.maxTime-cashier.minTime) + cashier.minTime
			time.Sleep(time.Duration(randTime) * time.Millisecond)
			fmt.Printf("car %d fuel: %d  paid the price \n", cashier.line[0].id, cashier.line[0].fuel)
			cashier.line[0].state = paid
			cashier.MoveInLine()
		}
	}

}
func (cashier *Cashier) CanAdd() bool {
	return len(cashier.line) < cashier.maxInLine
}
func (cashier *Cashier) MoveInLine() {
	cashier.mu.Lock()
	cashier.line = cashier.line[1:]
	cashier.mu.Unlock()
}
func (cashier *Cashier) AddToLine(car *car) bool {
	cashier.mu.Lock()
	if !cashier.CanAdd() {
		cashier.mu.Unlock()
		return false
	}
	car.state = waitingAtCashier
	cashier.line = append(cashier.line, car)
	cashier.mu.Unlock()
	return true
}

func CreateCashier(maxCarsInLine int, minTime int, maxTime int) *Cashier {
	c := new(Cashier)
	c.line = make([]*car, 0)
	c.minTime = minTime
	c.maxTime = maxTime
	c.maxInLine = maxCarsInLine
	return c

}
