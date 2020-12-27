package main

import "fmt"

func main() {
	jb := []string{"James", "Bond", "Shaken, not stirred"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Hello, James"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	for i, v := range xp { // if using "for _, v", the "_" means that we don't need the "index", so it will be thrown away
		fmt.Println("record:", i)
		for j, val := range v {
			fmt.Printf("\tindex position: %v\thas value: %v\n", j, val)
		}
	}
}
