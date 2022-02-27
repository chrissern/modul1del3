package main

import "testing"

func TestMoveBoat(t *testing.T) {
	if MoveBoat("West") != "East" {
		t.Error("Expected to move from West to East")
	} else if MoveBoat("East") != "West" {
		t.Error("Expected to move from East to West")
	}
}
