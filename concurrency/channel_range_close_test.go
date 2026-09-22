package concurrency

import (
	"reflect"
	"testing"
)

func TestChannelRangeClose(t *testing.T) {
	want := []string{
		"We", "use", "GO!",
	}

	got := getMessage(want)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}
