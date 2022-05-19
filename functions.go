package goessentials

import (
	"fmt"
	"http"
	"net/http"
)

func add(a int, b int) int {
	return a + b
}

func divmod(a int, b int) (int, int) {
	return a / b, a % b
}

func contentType(url string) (string, error) {
	resp, error := http.Get(url)

	if error != nil {
		return "", error
	}

	defer resp.Body.Close() //Make sure we close the body

	ctype := resp.Header.Get("Content-Type")

	if ctype == "" {
		return "", fmt.Errorf("Cant find content type header")
	}

	return ctype, nil
}

func main() {
	val := add(1, 2)

	fmt.Println(val)

	div, mod := divmod(7, 2)

	fmt.Printf("div=%d, mod=%d\n", div, mod)
}
