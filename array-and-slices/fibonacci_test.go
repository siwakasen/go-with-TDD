package arrayandslices

import (
	"slices"
	"testing"
)

// fibonacci numbers: 0,1,1,2,3,5,8,
func TestFibonacci(t *testing.T) {
	t.Run("fibonacci number 7", func(t *testing.T) {
		n := 7

		want := []int{0, 1, 1, 2, 3, 5, 8}
		got := fibbonacci(n)
		t.Log(got)
		if !slices.Equal(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("fibonacci number 1", func(t *testing.T) {
		n := 1

		want := []int{0}
		got := fibbonacci(n)

		if !slices.Equal(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("fibonacci number 2", func(t *testing.T) {
		n := 2

		want := []int{0, 1}
		got := fibbonacci(n)

		if !slices.Equal(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})
	t.Run("fibonacci number 3", func(t *testing.T) {
		n := 3

		want := []int{0, 1, 1}
		got := fibbonacci(n)

		if !slices.Equal(want, got) {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
