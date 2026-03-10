// Package stringmanipulation
package stringmanipulation

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"

	want := "olleh"

	got := reverseString(s)

	checkString(t, got, want)
}

func TestIsPalindrom(t *testing.T) {
	t.Run("palimdrom normal letter", func(t *testing.T) {
		original := "alla"
		want := original
		got := reverseString(original)

		checkString(t, got, want)
	})

	t.Run("palindrom with filtering non-letter", func(t *testing.T) {
		original := "al1l.A9#"

		want := "alla"
		got := palindromOnlyLetter(original)

		checkString(t, got, want)
	})
}

func checkString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
