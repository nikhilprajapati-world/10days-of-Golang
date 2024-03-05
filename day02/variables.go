package main

import "fmt"

func main() {
	var x int = 1
	fmt.Println("Number is:", x)

	var pi float64 = 3.1456
	fmt.Println(pi)

	var is_flag bool = true
	fmt.Print(is_flag)

	var uname string = "admin"
	var password string = "demo@12345"
	fmt.Println("User credentails are: ", uname+":"+password)

	//Also you can declare short variables.

	firstName := "John" // Initiates as string
	lastName := "Wick"  // same as above
	fmt.Println("User is: " + firstName + " " + lastName)

	// Example of int, string, bool.
	planets := 8
	planetname := "Earth"
	var isFar = true

	fmt.Println("Total number of planets are: ", planets)
	fmt.Println("Third planet name is: ", planetname)
	fmt.Println("Is Nepture far from Earth :", isFar)

	// Multiple strings
	make, model := "Tesla", "S1"
	fmt.Println("My car's make is: " + make + " and model is: " + model)

}
