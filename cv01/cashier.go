package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Cashier struct {
	line      []*car
	minTime   int `json:"min_time"`
	maxTime   int `json:"max_time"`
	mu        sync.Mutex
	maxInLine int `json:"max_in_line"`

	timeInLineTotal int64
	totalcars       int
}

func (cashier *Cashier) Print() {
	fmt.Printf("Cashier\n	total time in line: %d s\n	total cars: %d \n	avg. time in line %d ms\n",
		cashier.timeInLineTotal/1000, cashier.totalcars, cashier.timeInLineTotal/int64(cashier.totalcars))
}

func (cashier *Cashier) mainLoop() {
	for {
		if len(cashier.line) != 0 {
			var randTime = rand.Intn(cashier.maxTime-cashier.minTime) + cashier.minTime
			time.Sleep(time.Duration(randTime) * time.Millisecond)
			fmt.Printf("car %d fuel: %s paid the price \n", cashier.line[0].id, GetFuelName(cashier.line[0].fuel))
			cashier.line[0].state = paid
			cashier.timeInLineTotal += time.Now().UnixMilli() - cashier.line[0].cashierArrived
			cashier.totalcars++
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

func (c *Cashier) CreateCashier() {
	c.line = make([]*car, 0)
}
