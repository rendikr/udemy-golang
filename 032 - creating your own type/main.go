package main

import "fmt"

// declare a custom type
type hotdog int

// declare a variable to be of a certain type
var a int
var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
