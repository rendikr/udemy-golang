package main

import "fmt"

func main() {
	// x := type{values} // composite literal
	x := []int{4, 5, 7, 8, 42}
	fmt.Println(x)      // prints [4, 5, 7, 8, 42]
	fmt.Println(x[1])   // prints 5
	fmt.Println(x[1:])  // prints [5, 7, 8, 42]
	fmt.Println(x[1:3]) // prints [5, 7]

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i < len(x); i++ {
		fmt.Println(i, x[i])
	}
}

// a SLICE allows to group together VALUES of the same TYPE
