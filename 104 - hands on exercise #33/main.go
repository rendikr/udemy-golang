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

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	fmt.Println(m) // prints map[Bond:{James Bond [Chocolate Blueberry]} Moneypenny:{Miss Moneypenny [Vanilla Strawberry]}]

	for k, v := range m {
		fmt.Println(k)
		fmt.Println(v.first)
		fmt.Println(v.last)

		for i, val := range v.flavors {
			fmt.Println(i, val)
		}
		fmt.Println("----")
	}
}
