// Package arrayandslices contains fibbonacci function
package arrayandslices

func fibbonacci(n int) []int {
	if n == 1 {
		return []int{0}
	}
	fib := []int{0, 1}

	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	return fib[:n]
}
