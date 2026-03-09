package main

import (
	"fmt"
	"unicode"

	vrdc "github.com/siwakasen/go-with-TDD/variadic"
)

func main() {
	numbers := []int{1, 2, 3, 4}

	fmt.Printf("%d\n", vrdc.Calculate(numbers...))

	s := "hello, world"

	r := []rune(s)

	fmt.Printf("%q\n", r[0])

	fmt.Printf("%+v\n", unicode.IsLetter(r[0]))

	fmt.Printf("%+v\n", unicode.IsDigit(r[0]))
}
