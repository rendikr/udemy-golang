package main

import "fmt"

func main() {
	p1 := struct {
		first string
		last  string
		age   int
	}{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	fmt.Println(p1) // prints {James Bond 32}

	fmt.Println(p1.first, p1.last, p1.age) // prints James Bond 32
}
