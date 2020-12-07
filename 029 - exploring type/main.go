package main

import "fmt"

// declare a variable and assign a value to it.
// because it's outside the main function, it will be available globally
var y = 42
var z = "Shaken, not stirred"
var a = `James said, "Shaken, not stirred"`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y) // will print out the variable type of identifier 'y'
	fmt.Println(z)
	fmt.Printf("%T\n", z) // will print out the variable type of identifier 'z'
	fmt.Println(a)
	fmt.Printf("%T\n", a) // will print out the variable type of identifier 'a'
}
