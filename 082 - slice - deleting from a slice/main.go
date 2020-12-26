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

	x = append(z[:2], z[4:]...)
	// z[:2] will only take the z[0] & z[1]
	// z[4:] will take from z[4] to the end
	// this resulting in the value of [7, 8] to be taken out from z, and assign it as variable x
	fmt.Println(x) // prints [4 5 42 75 19 20 33 67]
}

// a SLICE allows to group together VALUES of the same TYPE
