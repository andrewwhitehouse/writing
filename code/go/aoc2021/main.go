package main

import (
	"fmt"
	"aoc2021/day1"
)

func main() {
	depths := []int16{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	fmt.Println(day1.CountIncreases(depths))
}

