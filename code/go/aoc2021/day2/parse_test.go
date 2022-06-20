package day2

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	input := "forward 5\ndown 5\nforward 8\nup 3\ndown 8"
	actual := Parse(input)
	expected := []*Command{
		newCommand(Forward, 5),
		newCommand(Down, 5),
		newCommand(Forward, 8),
		newCommand(Up, 3),
		newCommand(Down, 8),
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Parse was incorrect, got: %v, want: %v.", actual, expected)
	}
}
