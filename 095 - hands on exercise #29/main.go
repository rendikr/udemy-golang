package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james": []string{
			"Shaken, not stirred", "Martinis", "Women",
		},
		"moneypenny_miss": []string{
			"James Bond", "Literature", "Computer Science",
		},
		"chiffre_le": []string{
			"Cards", "Money", "Stock",
		},
	}

	for k, v := range m {
		fmt.Println("This is the record for", k)
		for i, val := range v {
			fmt.Println("\t", i, val)
		}
	}
}
