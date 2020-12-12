package main

import "fmt"

// declare a variable to be of a certain type
var x bool

func main() {
	fmt.Println(x)

	x = true

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	a := 7
	b := 42

	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a <= b)
	fmt.Println(a >= b)
}
