package main

import "fmt"

// declare a variable to be of a certain type
var x int
var y float64
var z int8

// uint8 : unsigned integer
// int8 : signed integer

func main() {
	x = 42
	y = 156.34534
	z = -128

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
