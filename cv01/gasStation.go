package main

type gasStation struct {
	pumps    []*Pump
	cashiers []*Cashier
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
func (gasStation *gasStation) GetBestCashier() *Cashier {
	var bestCashier *Cashier
	for _, cashier := range gasStation.cashiers {
		if bestCashier == nil || len(cashier.line) < len(bestCashier.line) {
			bestCashier = cashier
		}
	}
	return bestCashier
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
