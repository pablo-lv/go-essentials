package main

import "fmt"

func main() {
	// Same type
	loons := []string{"bugs", "daffy", "taz"}

	//length
	fmt.Println(len(loons))

	// 0 indexing
	fmt.Println(loons[0])

	//slices
	fmt.Print(loons[1:])

	//single value range
	for i := range loons {
		fmt.Println(i)
	}

	// double value range
	for i, name := range loons {
		fmt.Println("%s at %d\n", name, i)
	}

	// double value range, ignoring index by using _
	for _, name := range loons {
		fmt.Println(name)
	}

	//append
	loons = append(loons, "elmer")
}
