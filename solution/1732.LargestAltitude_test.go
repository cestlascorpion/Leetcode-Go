package solution

import "testing"

func TestLargestAltitude(t *testing.T) {
	gain := []int{-5, 1, 5, 0, -7}
	exp := 1
	if LargestAltitude(gain) != exp {
		t.Fail()
	}
}

func TestLargestAltitude2(t *testing.T) {
	gain := []int{-4, -3, -2, -1, 4, 3, 2}
	exp := 0
	if LargestAltitude(gain) != exp {
		t.Fail()
	}
}
