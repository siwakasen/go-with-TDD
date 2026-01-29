// Package arrayandslices is package that provides functions for array and slices test cases
package arrayandslices

func Sum(numbers []int) (sum int) {
	for _, v := range numbers {
		sum += v
	}

	return
}
