package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x) // prints [4, 5, 7, 8, 42]

	x = append(x, 75) // will add a value to the end of a slice
	fmt.Println(x)    // prints [4, 5, 7, 8, 42, 75]

	y := []int{19, 20, 33, 67}
	fmt.Println(y) // prints [19 20 33 67]

	z := append(x, y...)
	fmt.Println(z) // prints [4 5 7 8 42 75 19 20 33 67]
}

// a SLICE allows to group together VALUES of the same TYPE
