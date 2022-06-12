package day1

import "testing"

func TestCountIncreases(t *testing.T) {
	depths := []uint16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	var expected uint16 = 7
	actual := CountIncreases(depths)
	if expected != actual {
		t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
	}
}

func TestIntegration(t *testing.T) {
	actual, _ := Part1("199\n200\n208\n210\n200\n207\n240\n269\n260\n263")
	expected := uint16(7)
	if expected != actual {
		t.Errorf("Part1 was incorrect, got: %d, want: %d.", actual, expected)
	}
}
