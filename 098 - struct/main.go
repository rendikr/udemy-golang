package main

import "fmt"

// create a VALUE of certain TYPE to aggregate together VALUES of different types
type person struct {
	first string
	last  string
	age   int
}

func main() {
	// VALUE of a certain TYPE is equal to OBJECTS in other programming languages
	p1 := person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(p1) // prints {James Bond 32}
	fmt.Println(p2) // prints {Miss Moneypenny 27}

	fmt.Println(p1.first, p1.last, p1.age) // prints James Bond 32
	fmt.Println(p2.first, p2.last, p2.age) // prints Miss Moneypenny 27
}
