package main

import "fmt"

func main() {
	// String are immutable
	book := "The colour of the magic"

	fmt.Println(len(book))

	//slice(start, end), 0 based, 1/2 empty range
	fmt.Println(book[4:11])

	//slice (no end)
	fmt.Println(book[4:])

	//slice (no start)
	fmt.Println(book[:4])

	// use + to concatenate strings
	fmt.Println("t" + book)

	//multi line
	multiline := `
	This is the first line
	This is the second line
	...
	`

	fmt.Println(multiline)
}
