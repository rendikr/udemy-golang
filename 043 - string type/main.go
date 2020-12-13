package main

import "fmt"

// declare a variable to be of a certain type


func main() {
	s := "Hello, github"
	fmt.Println(s)

	s = `"Hello, github"`
	fmt.Println(s)

	fmt.Printf("%T\n", s)

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
