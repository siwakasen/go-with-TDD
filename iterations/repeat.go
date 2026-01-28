// Package iterations is package that provides functions for iteration test cases
package iterations

import "strings"

func Repeat(character string) string {
	var repeated strings.Builder

	for range 5 {
		repeated.WriteString(character)
	}

	return repeated.String()
}

func OldRepeat(character string) string {
	var repeated string

	for range 5 {
		repeated += character
	}

	return repeated
}
