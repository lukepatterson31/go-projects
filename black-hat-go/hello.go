package main

import (
	"fmt"
)

func main() {
	var count = int(42)

	// can also be declared
	// ptr := &count
	var ptr *int = &count

	fmt.Println(*ptr)

	*ptr = int(100)

	fmt.Println(*ptr)
	fmt.Println(count)
	fmt.Println("Hello, Black Hat Gophers!")
}
