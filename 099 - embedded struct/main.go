package main

import "fmt"

// create a VALUE of certain TYPE to aggregate together VALUES of different types
type person struct {
	first string
	last  string
	age   int
}

type secretAgent struct {
	person
	licenseToKill bool
}

func main() {
	// VALUE of a certain TYPE is equal to OBJECTS in other programming languages
	sa1 := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		licenseToKill: true,
	}

	p2 := person{
		first: "Miss",
		last:  "Moneypenny",
		age:   27,
	}

	fmt.Println(sa1) // prints {{James Bond 32} true}
	fmt.Println(p2)  // prints {Miss Moneypenny 27}

	fmt.Println(sa1.first, sa1.person.last, sa1.age, sa1.licenseToKill) // prints James Bond 32 true
	// child struct can be called on the parent (sa1.first, while "first" is a value on the child struct)
	fmt.Println(p2.first, p2.last, p2.age) // prints Miss Moneypenny 27
}
