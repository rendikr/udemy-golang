package main

import "fmt"

func main() {
	x := struct {
		first     string
		friends   map[string]int
		favColors []string
	}{
		first: "James",
		friends: map[string]int{
			"Moneypenny": 555,
			"Q":          888,
			"M":          999,
		},
		favColors: []string{"red", "gray"},
	}

	fmt.Println(x) // prints {James map[M:999 Moneypenny:555 Q:888] [red gray]}

	for i, v := range x.friends {
		fmt.Println(i, v)
	}

	for i, v := range x.favColors {
		fmt.Println(i, v)
	}
}
