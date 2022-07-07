package day2

import (
	"reflect"
	"testing"
)

func TestNavigation(t *testing.T) {
	commands := []Command{
		newCommand(Forward, 5),
		newCommand(Down, 5),
		newCommand(Forward, 8),
		newCommand(Up, 3),
		newCommand(Down, 8),
		newCommand(Forward, 2),
	}
	position := Navigate(commands)
	expected := Position{distance: 15, depth: 10}
	if !reflect.DeepEqual(position, expected) {
		t.Errorf("End position was incorrect, got: %v, want: %v.", position, expected)
	}
}

func TestPart1Integration(t *testing.T) {
	expected := uint64(150)
	actual, _ := Part1("forward 5\ndown 5\nforward 8\nup 3\ndown 8\nforward 2")
	if actual != expected {
		t.Errorf("Part1 result was incorrect, got: %d, want: %d.", actual, expected)
	}
}
