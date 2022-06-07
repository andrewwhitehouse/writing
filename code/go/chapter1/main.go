package main

import "fmt"

func Sum(x int, y int) int {
    return x + y
}

func main() {
    var x int = 5
    var y int = 10
    fmt.Printf("Sum of %d and %d is %d\n", x, y, Sum(x, y))
}
