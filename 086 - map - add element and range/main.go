package main

import "fmt"

func main() {
	m := map[string]int{
		"James":      32,
		"Moneypenny": 27,
	}

	fmt.Println(m)          // prints map[James:32 Moneypenny:27]
	fmt.Println(m["James"]) // prints 32
	fmt.Println(m["Q"])     // prints 0

	v, ok := m["Q"]
	fmt.Println(v)  // prints 0
	fmt.Println(ok) // prints false, because "Q" is not existed on the map

	m["Q"] = 26 // adds a new key value pair to the map

	if v, ok := m["Q"]; ok {
		fmt.Println("Q found!", v) // prints the value of m["Q"] if it's existed
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

}
