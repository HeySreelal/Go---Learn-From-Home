package main

import (
	"fmt"
)

func main() {
	var x, y int = 3, 4
	var a, b = float64(x), float64(y)
	fmt.Printf("X: %v - %T\nY: %v - %T\nA: %v - %T\nB: %v - %T\n", x, x, y, y, a, a, b, b)
}
