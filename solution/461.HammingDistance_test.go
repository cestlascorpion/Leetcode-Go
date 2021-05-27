package solution

import "testing"

func TestHammingDistance(t *testing.T) {
	x, y := 1, 4
	exp := 2
	if exp != HammingDistance(x, y) {
		t.Fail()
	}
}
