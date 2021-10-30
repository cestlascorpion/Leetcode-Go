package solution

import "testing"

func TestFindMiddleIndex(t *testing.T) {
	nums := []int{2, 3, -1, 8, 4}
	exp := 3
	if FindMiddleIndex(nums) != exp {
		t.FailNow()
	}
}

func TestFindMiddleIndex2(t *testing.T) {
	nums := []int{1, -1, 4}
	exp := 2
	if FindMiddleIndex(nums) != exp {
		t.FailNow()
	}
}

func TestFindMiddleIndex3(t *testing.T) {
	nums := []int{2, 5}
	exp := -1
	if FindMiddleIndex(nums) != exp {
		t.FailNow()
	}
}

func TestFindMiddleIndex4(t *testing.T) {
	nums := []int{1}
	exp := 0
	if FindMiddleIndex(nums) != exp {
		t.FailNow()
	}
}
