package states

import "fmt"

var FarmerPosition = ""
var ChickenPosition = ""
var FoxPosition = ""
var CornPosition = ""
var BoatPosition = ""

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
	Position()
	Test()
}

func Position() {
	if FarmerPosition == "West" {
		westSlot("Farmer")
	} else if FarmerPosition == "Boat" {
		boatSlot("Farmer")
	} else if FarmerPosition == "East" {
		eastSlot("Farmer")
	}

	if ChickenPosition == "West" {
		westSlot("Chicken")
	} else if ChickenPosition == "Boat" {
		boatSlot("Chicken")
	} else if ChickenPosition == "East" {
		eastSlot("Chicken")
	}

	if FoxPosition == "West" {
		westSlot("Fox")
	} else if FoxPosition == "Boat" {
		boatSlot("Fox")
	} else if FoxPosition == "East" {
		eastSlot("Fox")
	}

	if CornPosition == "West" {
		westSlot("Corn")
	} else if CornPosition == "Boat" {
		boatSlot("Corn")
	} else if CornPosition == "East" {
		eastSlot("Corn")
	}
}

func PlaceBoat() {
	boat = "\\_" + bSlot1 + "_" + bSlot2 + "_/"
	if BoatPosition == "West" {
		wCoast = boat
		eCoast = ""
	} else if BoatPosition == "East" {
		eCoast = boat
		wCoast = ""
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
}

func Test() {
	worldSituation = "_" + wSlot1 + "_" + wSlot2 + "_" + wSlot3 + "_" + wSlot4 + "_W|" + wCoast + "_________" + eCoast + "|E_" + eSlot1 + "_" + eSlot2 + "_" + eSlot3 + "_" + eSlot4 + "_"
	fmt.Print(worldSituation)
}
