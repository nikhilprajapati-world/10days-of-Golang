package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("The current system Time is: ", currentTime)

	fmt.Print("\nCurrent Month is: ", currentTime.Month())
	fmt.Print("\nCurrent Date is: ", currentTime.Day())
	fmt.Print("\nCurrent Year is: ", currentTime.Year())
	fmt.Print("\nCurrent Weekday is: ", currentTime.Weekday())

	fmt.Println("\n")
	fmt.Println("Example of custom Date in dd-mm-yyyy hh:mm:ss:")

	// Prints custom date and time in below format
	fmt.Printf("%d-%d-%d %d:%d:%d\n",
		currentTime.Day(),
		currentTime.Month(),
		currentTime.Year(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second())

}
