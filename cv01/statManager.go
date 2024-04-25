package main

import (
	"fmt"
	"sync"
	"time"
)

type StatManager struct {
	totalCarAtPumpTimeLock   sync.Mutex
	totalCarAtGasStationTime int64
	maxCarAtPumpTimeLock     sync.Mutex
	maxCarAtGasStationTime   int64
}

var singleton *StatManager = &StatManager{}

func GetStatManager() *StatManager {
	return singleton
}
func (this *StatManager) CarLeft(carCreation int64) {
	var cartime = time.Now().UnixMilli() - carCreation
	this.totalCarAtPumpTimeLock.Lock()
	this.totalCarAtGasStationTime += cartime
	this.totalCarAtPumpTimeLock.Unlock()
	if cartime > this.maxCarAtGasStationTime {
		this.maxCarAtPumpTimeLock.Lock()
		this.maxCarAtGasStationTime = cartime
		this.maxCarAtPumpTimeLock.Unlock()
	}
}
func (this *StatManager) Print(totalcars int) {
	fmt.Printf("total cars: %d\ntotal time: %d s\n avg time: %d ms\n max time: %d ms\n",
		totalcars, this.totalCarAtGasStationTime/1000, this.totalCarAtGasStationTime/int64(totalcars), this.maxCarAtGasStationTime)
}
