package main

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

func GetFuelName(fuel int) string {
	switch fuel {
	case gas:
		return "gas"
	case diesel:
		return "diesel"
	case LPG:
		return "lPG"
	case electric:
		return "electric"
	case numberOfFuel:
		return "numberOfFuel"
	}
	return "unknown"
}
