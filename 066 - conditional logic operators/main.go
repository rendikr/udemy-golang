package main

import "fmt"

func main() {
	fmt.Print("true && true\t %v\n", true && true)
	fmt.Print("true && false\t %v\n", true && false)
	fmt.Print("true || true\t %v\n", true || true)
	fmt.Print("true || false\t %v\n", true || false)
	fmt.Print("!true\t %v\n", !true)

	fmt.Println("done.")
}
