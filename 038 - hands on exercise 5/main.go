package main

import "fmt"

// declare a variable to be of a certain type
type family int
var x family
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
	fmt.Printf("%v\n", x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
}
