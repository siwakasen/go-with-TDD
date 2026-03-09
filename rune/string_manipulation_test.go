// Package stringmanipulation
package stringmanipulation

import (
	"testing"
)

func TestReverseString(t *testing.T) {
	s := "hello"

	want := "olleh"

	got := reverseString(s)

	if want != got {
		t.Errorf("got %s want %s", got, want)
	}
}
