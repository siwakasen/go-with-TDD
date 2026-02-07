package stringtochar

import (
	"slices"
	"testing"
)

func TestStringToChar(t *testing.T) {
	got := StringToChar("Hello, World!")
	want := []string{"H", "e", "l", "l", "o", ",", " ", "W", "o", "r", "l", "d", "!"}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
