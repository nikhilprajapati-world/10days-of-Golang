package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// Seed is used to genearte different value every time
	rand.Seed(time.Now().UnixNano())

	// Intn create randon number as per n value
	fmt.Println(rand.Intn(200)) // random no from 0 to 200
	fmt.Println(rand.Intn(10))  // random no from 0 to 10
	fmt.Println(rand.Intn(5))   // random no from 0 to 5
	fmt.Println()

}
