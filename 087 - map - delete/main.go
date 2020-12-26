package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m) // prints map[James:32 Moneypenny:27]

	delete(m, "James")

	fmt.Println(m) // prints map[Moneypenny:27]

	if v, ok := m["Moneypenny"]; ok {
		fmt.Println("value:", v) // prints 27
		delete(m, "Moneypenny")
	}

	fmt.Println(m) // prints map[]
}
