package main

import (
	"fmt"
)

// Addition simple function
func add(x, y int) int {
	return x + y

}

// Substraction simple function
func sub(a, b int) int {
	return a - b
}

func main() {
	fmt.Println("Addition example: ", add(1, 4))
	fmt.Println("Substraction example: ", sub(50, 23))
}
