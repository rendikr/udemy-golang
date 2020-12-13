package main

import "fmt"

const (
	// kb = 1024
	// mb = 1024 * kb
	// gb = 1024 * mb

	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	x := 2
	fmt.Printf("%d\t\t%b\n", x, x)

	// bit shifting the binary from x
	// before: 0000 0010
	y := x << 1
	// after: 0000 0100
	fmt.Printf("%d\t\t%b\n", y, y)

	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}
