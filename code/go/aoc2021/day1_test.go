package day1

import "testing"

func TestSum(t *testing.T) {
    depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
    var expected uint16 = 7
    actual := CountIncreases(depths)
    if expected != actual {
       t.Errorf("CountIncreases was incorrect, got: %d, want: %d.", actual, expected)
    }
}
