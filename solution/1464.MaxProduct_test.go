package solution

import "testing"

func TestMaxProduct(t *testing.T) {
	nums := []int{3, 4, 5, 2}
	exp := 12
	if MaxProduct(nums) != exp {
		t.Fail()
	}
}

func TestMaxProduct2(t *testing.T) {
	nums := []int{3, 7}
	exp := 12
	if MaxProduct(nums) != exp {
		t.Fail()
	}
}
