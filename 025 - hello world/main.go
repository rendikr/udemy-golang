package main

import "fmt"

func main() {
	fmt.Println("Hello Github")
	foo()
	fmt.Println("Even Number :")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
}

func foo() {
	fmt.Println("bar")
}

func bar() {
	fmt.Println("baz")
}
