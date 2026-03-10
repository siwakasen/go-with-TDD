// Package stringmanipulation
package stringmanipulation

import (
	"slices"
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"

	want := "olleh"

	got := reverseString(s)

	checkTwoSrings(t, got, want)
}

func TestIsPalindrom(t *testing.T) {
	t.Run("palimdrom normal letter", func(t *testing.T) {
		original := "alla"
		want := original
		got := reverseString(original)

		checkTwoSrings(t, got, want)
	})

	t.Run("palindrom with filtering non-letter", func(t *testing.T) {
		original := "al1l.A9#"

		want := "alla"
		got := palindromOnlyLetter(original)

		checkTwoSrings(t, got, want)
	})
}

func TestStringChecker(t *testing.T) {
	t.Run("is digit?", func(t *testing.T) {
		s := "1"
		want := true
		got := isDigit(s[0])

		checkTwoBoolean(t, got, want)
	})

	t.Run("is Symbol", func(t *testing.T) {
		s := "%.*#;,"
		want := []bool{true, true, true, true, true, true}

		got := func() []bool {
			var result []bool

			for i := 0; i < len(s); i++ {
				if isSymbol(s[i]) {
					result = append(result, true)
				} else {
					result = append(result, true)
				}
			}
			return result
		}()

		if !slices.Equal(got, want) {
			t.Errorf("got %+v want %+v", got, want)
		}
	})
}

func checkTwoBoolean(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("got %t, want %t", got, want)
	}
}

func checkTwoSrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
