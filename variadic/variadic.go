package variadic

func Calculate(numbers ...int) int {
	sum := 0
	for _, v := range numbers {
		sum += v
	}

	return sum
}
