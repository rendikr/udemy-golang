package main

import "fmt"

// declare a variable to be of a certain type
var y string
var z int

func main() {
	fmt.Println("printing the value of y", y, ". printing completed")
	fmt.Printf("%T\n", y)

	// assign a value of that type to the variable
	y = "Bond, James Bond"

	fmt.Println("printing the value of y", y, ". printing completed")
	fmt.Printf("%T\n", y)

	fmt.Println("printing the value of z", z, ". printing completed")
	fmt.Printf("%T\n", z)
}
