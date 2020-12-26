package main

import "fmt"

func main() {
	x := make([]int, 10, 12) // make([]T, length, capacity)
	fmt.Println(x)           // prints [0,0,0,0,0,0,0,0,0,0]
	fmt.Println(len(x))      // prints 10
	fmt.Println(cap(x))      // prints 12

	x[0] = 42
	x[9] = 999

	fmt.Println(x)      // prints [42,0,0,0,0,0,0,0,0,999]
	fmt.Println(len(x)) // prints 10
	fmt.Println(cap(x)) // prints 12

	x = append(x, 3423)

	fmt.Println(x)      // prints [42,0,0,0,0,0,0,0,0,999, 3423]
	fmt.Println(len(x)) // prints 11
	fmt.Println(cap(x)) // prints 12

	x = append(x, 8179)

	fmt.Println(x)      // prints [42,0,0,0,0,0,0,0,0,999, 3423, 8179]
	fmt.Println(len(x)) // prints 12
	fmt.Println(cap(x)) // prints 12

	x = append(x, 8667)

	fmt.Println(x)      // prints [42,0,0,0,0,0,0,0,0,999, 3423, 8179, 8667]
	fmt.Println(len(x)) // prints 13
	fmt.Println(cap(x)) // prints 24
}
