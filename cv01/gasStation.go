package main

type gasStation struct {
	Pumps    []*Pump    `json:"pumps"`
	Cashiers []*Cashier `json:"cashiers"`
}

func (gasStation *gasStation) GetBestPump(fuel int) *Pump {
	var bestPump *Pump
	for _, pump := range gasStation.Pumps {
		if pump.Fuel != fuel {
			continue
		}
		if bestPump == nil || len(pump.line) < len(bestPump.line) {
			bestPump = pump
		}
	}
	return bestPump
}
func (gasStation *gasStation) GetBestCashier() *Cashier {
	var bestCashier *Cashier
	for _, cashier := range gasStation.Cashiers {
		if bestCashier == nil || len(cashier.line) < len(bestCashier.line) {
			bestCashier = cashier
		}
	}
	return bestCashier
}

func (gasStation *gasStation) Open() {
	for _, pump := range gasStation.Pumps {
		pump.CreatePump()
		go pump.mainLoop(gasStation)
	}
	for _, cashier := range gasStation.Cashiers {
		cashier.CreateCashier()
		go cashier.mainLoop()
	}
}
func (gasStation *gasStation) EveryOneIsGone() bool {
	for _, pump := range gasStation.Pumps {
		if len(pump.line) != 0 {
			return false
		}
	}
	return true
}
func (gasStation *gasStation) Print() {
	for _, pump := range gasStation.Pumps {
		pump.Print()
	}
	for _, cashier := range gasStation.Cashiers {
		cashier.Print()
	}
}
