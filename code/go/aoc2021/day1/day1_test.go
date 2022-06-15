package day1

import (
	"reflect"
	"testing"
)

func TestCountIncreases(t *testing.T) {
	depths := []uint16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	var expected uint16 = 7
	actual := CountIncreases(depths)
	if expected != actual {
		t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart1Integration(t *testing.T) {
	actual, _ := Part1("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(7)
	if expected != actual {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestPart2Integration(t *testing.T) {
	actual, _ := Part2("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(5)
	if expected != actual {
		t.Errorf("Part2 was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestSlidingWindow(t *testing.T) {
	depths := []uint16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := []uint16{607, 618, 618, 617, 647, 716, 769, 792}
	actual := SlidingWindow(depths)
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("SlidingWindow was incorrect, got: %d, want: %d.", actual, expected)
	}
}
