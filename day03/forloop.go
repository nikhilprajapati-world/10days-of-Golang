package main

import "fmt"

func main() {

	// Prints 1 to 5
	fmt.Println("Print 1 to 5 in forloop: ")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}

	// same way in less lines of code
	fmt.Println("\nAnother way to print in loop: ")

	for x := 1; x <= 5; x++ {
		fmt.Println(x)
	}

	// Prints blank line
	fmt.Println("")

	planetno := 1
	for planetno < 9 {
		fmt.Println("Planet count: ", planetno)
		planetno++
	}

}
