package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("false should not be printed")
	case true:
		fmt.Println("true should be printed")
	}

}
