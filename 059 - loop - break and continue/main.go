package main

import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40

	fmt.Println(x, y)

	i := 1
	for {
		if i > 100 {
			break
		}

		if i%2 != 0 {
			i++
			continue
		}

		fmt.Println(i)
		i++
	}
	fmt.Println("done.")
}
