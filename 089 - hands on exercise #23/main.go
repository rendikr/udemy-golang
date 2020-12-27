package main

import "fmt"

func main() {
	x := []int{4, 5, 7, 8, 42, 55, 67, 81, 92, 99}
	fmt.Println(x)

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}
