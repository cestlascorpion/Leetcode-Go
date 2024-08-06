package solution

import "testing"

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}

func TestIsPowerOfTwo(t *testing.T) {
	n := 1
	if isPowerOfTwo(n) {
		t.Log("231. Power of Two: case 1 succeeded.")
	} else {
		t.Error("231. Power of Two: case 1 failed.")
	}
	n = 16
	if isPowerOfTwo(n) {
		t.Log("231. Power of Two: case 2 succeeded.")
	} else {
		t.Error("231. Power of Two: case 2 failed.")
	}
	n = 218
	if !isPowerOfTwo(n) {
		t.Log("231. Power of Two: case 3 succeeded.")
	} else {
		t.Error("231. Power of Two: case 3 failed.")
	}
}