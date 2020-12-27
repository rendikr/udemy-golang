package main

import "fmt"

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	car1 := truck{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		fourWheel: true,
	}

	car2 := sedan{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		luxury: true,
	}

	fmt.Println(car1) // prints {{4 black} true}
	fmt.Println(car2) // prints {{2 white} true}
}
