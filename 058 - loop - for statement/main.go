package main

import "fmt"

func main() {
	x := 1

	// for statement
	for x < 10 {
		fmt.Println(x)
		x++
	}
	fmt.Println("done.")

	// for clause
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	y := 1
	for {
		if y > 9 {
			break
		}
		fmt.Println(y)
		y++
	}
}
