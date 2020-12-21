package main

import "fmt"

func main() {
	favSport := "cycling"

	switch favSport {
	case "swimming":
		fmt.Println("you like to swim, go to the pool")
	case "hiking":
		fmt.Println("you like to hike, go to the mountain")
	case "cycling":
		fmt.Println("you like to cycle, go to the smooth road")
	}

}
