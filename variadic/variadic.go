// Package variadic
package variadic

func calculate(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}

func Calculate(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}
