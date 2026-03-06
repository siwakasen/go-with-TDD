package main

import (
	"fmt"

	"github.com/siwakasen/go-with-TDD/variadic"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Printf("%d", variadic.Calculate(numbers...))
}
