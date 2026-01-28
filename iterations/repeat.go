// Package iterations is package that provides functions for iteration test cases
package iterations

import (
	"strings"
)

func Repeat(character string, iterations int) string {
	var repeated strings.Builder

	for range iterations {
		repeated.WriteString(character)
	}

	return repeated.String()
}

func OldRepeat(character string, iterations int) string {
	var repeated string

	for range iterations {
		repeated += character
	}

	return repeated
}
