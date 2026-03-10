// Package stringmanipulation
package stringmanipulation

import "strings"

func isLetter(c byte) bool {
	return c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z'
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return !isLetter(c) && !isDigit(c) && c != ' ' && c != '\n' && c != '\t'
}

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func palindromOnlyLetter(s string) string {
	var letter strings.Builder

	for i := 0; i < len(s); i++ {
		if isLetter(s[i]) {
			letter.WriteByte(s[i])
		}
	}
	return reverseString(strings.ToLower(letter.String()))
}
