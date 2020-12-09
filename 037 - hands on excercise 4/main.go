package main

import "fmt"

// declare a variable to be of a certain type
type family int
var x family

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = 42
	fmt.Println(x)
	fmt.Printf("%v\n", x)
}
