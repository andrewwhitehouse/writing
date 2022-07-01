package day2

import (
	"testing"
)

func TestNewCommand(t *testing.T) {
	command := newCommand(Up, 20)
	actualDirection := command.direction
	expectedDirection := Up
	if actualDirection != expectedDirection {
		t.Errorf("Command direction was incorrect, got: %d, want: %d.", actualDirection, expectedDirection)
	}
	expectedDistance := uint16(20)
	actualDistance := command.distance
	if actualDistance != expectedDistance {
		t.Errorf("Command distance was incorrect, got: %d, want: %d.", actualDirection, expectedDirection)
	}
}
