package main

import (
	"fmt"

	vrdc "github.com/siwakasen/go-with-TDD/variadic"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	// using other package's function
	fmt.Printf("%d\n", vrdc.Calculate(numbers...))
}
