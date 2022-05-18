package main

import "fmt"
import "math"

func square(x float64) float64 {
	return x * x
}

func longestWallLength(x, y uint32) uint32 {
	return uint32(math.Sqrt(float64(x*x + y*y)))
}

func multiply1(x, y float64) float64 {
	return x * y
}

func multiple2(x float64, y float64) float64 {
	return x * y
}

func main() {
	//fmt.Printf("Max float64 %f\n", math.MaxFloat64);
	//fmt.Println(square(math.MaxInt32));
	fmt.Printf("Longest wall %dmm\n", longestWallLength(10*1000, 4950))
}
