package multiplereturnfunctions

import "testing"

func TestGetMinMax(t *testing.T) {
	wantMin, wantMax := -4, 12

	gotMin, gotMax := getMinMax([]int{4, -2, 1, 6, 9, 7, 12, 5, -4})

	if wantMin != gotMin && wantMax != gotMax {
		t.Errorf("got min %d want min %d, got max %d want max %d", gotMin, wantMin, gotMax, wantMax)
	}
}
