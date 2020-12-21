package main

import "fmt"

func main() {
	x := "James Bond"

	if x == "Quartermaster" {
		fmt.Printf("is Quartermaster?\t%v", x)
	} else if x == "James Bond" {
		fmt.Printf("is James Bond?\t%v", x)
	} else {
		fmt.Printf("neither\t%v", x)
	}

}
