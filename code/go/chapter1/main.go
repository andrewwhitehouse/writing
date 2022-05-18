package main

import "fmt"

func identity(x int32) int32 {
	return x
}

func main() {
	var age int64 = 22
	fmt.Println(identity(age))
}
