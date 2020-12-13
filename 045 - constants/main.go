package main

import "fmt"

const a = 42
const b = 22.78
const c = "James Bond"

// another approach to declare multiple constants
const (
	x = 42           // also can be written as x int = 42
	y = 22.78        // also can be written as y float64 = 22.78
	z = "James Bond" // also can be written as z string = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)
}
