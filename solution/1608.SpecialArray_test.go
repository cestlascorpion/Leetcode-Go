package solution

import "testing"

func TestSpecialArray(t *testing.T) {
	nums := []int{3, 5}
	exp := 2
	if SpecialArray(nums) != exp {
		t.Fail()
	}
}

func TestSpecialArray2(t *testing.T) {
	nums := []int{0, 4, 3, 0, 4}
	exp := 3
	if SpecialArray(nums) != exp {
		t.Fail()
	}
}

func TestSpecialArray3(t *testing.T) {
	nums := []int{3, 6, 7, 7, 0}
	exp := -1
	if SpecialArray(nums) != exp {
		t.Fail()
	}
}

func TestSpecialArray4(t *testing.T) {
	nums := []int{0, 0}
	exp := -1
	if SpecialArray(nums) != exp {
		t.Fail()
	}
}
