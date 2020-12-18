package main

import "fmt"

func main() {
	// A fallthrough statement transfers control to the next case.

	switch {
	case false:
		fmt.Println("this should not be printed")
	case (2 == 4):
		fmt.Println("this should not be printed either")
	case (3 == 3):
		fmt.Println("this should be printed")
		fallthrough
	case (4 == 4):
		fmt.Println("also true, but does it print?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true, should not be printed")
		fallthrough
	case (11 == 14):
		fmt.Println("not true, should not be printed either")
		fallthrough
	case (15 == 15):
		fmt.Println("true, should be printed")
		fallthrough
	default:
		fmt.Println("this is default")
	}

	name := "Bond"
	switch name {
	case "Moneypenny", "Bond", "Tanner":
		fmt.Println("miss money or bond or tanner")
	case "M":
		fmt.Println("m")
	case "Q":
		fmt.Println("q")
	default:
		fmt.Println("this is default")

	}

	fmt.Println("done.")
}
