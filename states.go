package main

import "fmt"

var farmerPosition string = ""
var chickenPosition string = ""
var foxPosition string = ""
var cornPosition string = ""
var boatPosition string = ""

var boat string = ""

var wSlot1 = ""
var wSlot2 = ""
var wSlot3 = ""
var wSlot4 = ""
var bSlot1 = ""
var bSlot2 = ""
var eSlot1 = ""
var eSlot2 = ""
var eSlot3 = ""
var eSlot4 = ""
var wCoast = ""
var eCoast = ""

var worldSituation = ""

func main() {
	farmerPosition = "Boat"
	chickenPosition = "West"
	foxPosition = "Boat"
	cornPosition = "West"
	boatPosition = "East"
	position()
	test()
}

func position() {
	if farmerPosition == "West" {
		westSlot("Farmer")
	} else if farmerPosition == "Boat" {
		boatSlot("Farmer")
	} else if farmerPosition == "East" {
		eastSlot("Farmer")
	}

	if chickenPosition == "West" {
		westSlot("Chicken")
	} else if chickenPosition == "Boat" {
		boatSlot("Chicken")
	} else if chickenPosition == "East" {
		eastSlot("Chicken")
	}

	if foxPosition == "West" {
		westSlot("Fox")
	} else if foxPosition == "Boat" {
		boatSlot("Fox")
	} else if foxPosition == "East" {
		eastSlot("Fox")
	}

	if cornPosition == "West" {
		westSlot("Corn")
	} else if cornPosition == "Boat" {
		boatSlot("Corn")
	} else if cornPosition == "East" {
		eastSlot("Corn")
	}

	if boatPosition == "West" {
		wCoast = boat
	} else if boatPosition == "East" {
		eCoast = boat
	}
}

func westSlot(item string) {
	if wSlot1 == "" {
		wSlot1 = item
	} else if wSlot2 == "" {
		wSlot2 = item
	} else if wSlot3 == "" {
		wSlot3 = item
	} else if wSlot4 == "" {
		wSlot4 = item
	}

}

func eastSlot(item string) {
	if eSlot1 == "" {
		eSlot1 = item
	} else if eSlot2 == "" {
		eSlot2 = item
	} else if eSlot3 == "" {
		eSlot3 = item
	} else if eSlot4 == "" {
		eSlot4 = item
	}
}

func boatSlot(item string) {
	if bSlot1 == "" {
		bSlot1 = item
	} else if bSlot2 == "" {
		bSlot2 = item
	}
	boat = "\\_" + bSlot1 + "_" + bSlot2 + "_/"
}

func test() {
	worldSituation = "_" + wSlot1 + "_" + wSlot2 + "_" + wSlot3 + "_" + wSlot4 + "_W|" + wCoast + "_________" + eCoast + "|E_" + eSlot1 + "_" + eSlot2 + "_" + eSlot3 + "_" + eSlot4 + "_"
	fmt.Print(worldSituation)
}
