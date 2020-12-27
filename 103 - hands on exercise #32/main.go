package main

import "fmt"

type person struct {
	first   string
	last    string
	flavors []string
}

func main() {
	p1 := person{
		first:   "James",
		last:    "Bond",
		flavors: []string{"Chocolate", "Blueberry"},
	}

	p2 := person{
		first:   "Miss",
		last:    "Moneypenny",
		flavors: []string{"Vanilla", "Strawberry"},
	}

	fmt.Println(p1) // prints {James Bond [Chocolate Blueberry]}
	fmt.Println(p2) // prints {Miss Moneypenny [Vanilla Strawberry]}

	for i, v := range p1.flavors {
		fmt.Println(i, v)
	}

	for i, v := range p2.flavors {
		fmt.Println(i, v)
	}
}
