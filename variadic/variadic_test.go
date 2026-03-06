package variadic

import "testing"

func TestVariadicFunction(t *testing.T) {
	numbers := []int{1, 2, 3, 4}

	got := Calculate(numbers...)
	want := 10

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
