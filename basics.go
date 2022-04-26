package main

import "fmt"

func main() {
	x := 10

	if x > 5 {
		fmt.Println("x is big")
	} else {
		fmt.Println("x is not that big")
	}

	if x > 5 && x < 15 {
		fmt.Println("x is just right")
	}

	a := 11.0
	b := 20.0

	if frac := a / b; frac > 0.5 {
		fmt.Println("a is more thatn half of b")
	}

	switch x {
	case 1:
		fmt.Println("one")
	default:
		fmt.Println("default")
	}

}
