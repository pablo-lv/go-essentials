package main

import "fmt"

func main() {
	_stocks := map[string]float64{
		"AMZ":  1699.8,
		"GOOG": 1129.19, // must have a multiline comma in multi line
	}

	//number of items
	fmt.Println(len(_stocks))
	// get a value
	fmt.Println(_stocks["AMZ"])
	// get zero value if not found
	fmt.Println(_stocks["TSLA"])
	// use two value to see if found
	value, ok := _stocks["TSLA"]

	if !ok {
		fmt.Println("TSLA not found")
	} else {
		fmt.Println(value)
	}

	//set
	_stocks["TSLA"] = 322.12

	//delete
	delete(_stocks, "AMZN")

}
