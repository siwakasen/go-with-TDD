// Package stringtochar is package that provides functions for array and slices test cases to sting
package stringtochar

func StringToChar(str string) []string {
	var sliceOfChar []string
	for _, r := range str {
		sliceOfChar = append(sliceOfChar, string(r))
	}

	return sliceOfChar
}
