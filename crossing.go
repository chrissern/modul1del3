package main

import (
	s "github.com/chrissern/modul1del3/states"
)

func main() {
	initialize()
}

func initialize() {
	s.FarmerPosition = "Boat"
	s.ChickenPosition = "West"
	s.FoxPosition = "Boat"
	s.CornPosition = "West"
	s.BoatPosition = "East"
	s.Position()
	s.PlaceBoat()
	s.Test()
}

func MoveBoat(item string) (result string) {
	if item == "West" {
		result = "East"
		s.BoatPosition = "East"
		s.PlaceBoat()
	} else if item == "East" {
		result = "West"
		s.BoatPosition = "West"
		s.PlaceBoat()
	}
	return result
}
