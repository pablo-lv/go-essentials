package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i > 2 {
			break
		}
	}

	for i := 0; i < 3; i++ {
		fmt.Println(i)
		if i < 1 {
			continue
		}
	}

	a := 0
	for a < 3 {
		fmt.Println(a)
		a++
	}

	// while
	b := 0
	for {
		if b > 2 {
			break
		}
		fmt.Println(b)
		b++
	}

}
