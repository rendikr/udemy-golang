package main

import "fmt"

// declare a variable and assign a value to it.
// because it's outside the main function, it will be available globally
var y = 43

// declare a variable with the identifier 'z'
// and that the variable is of type INT
// assigns a ZERO value to the variable (by default)
var z int

func main() {
	// short declaration operator ":=" is to declare & asssign value to a variable. only available inside a function body
	x := 42
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
