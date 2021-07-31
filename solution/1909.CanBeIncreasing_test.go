package solution

import "testing"

func TestCanBeIncreasing(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	if !CanBeIncreasing(nums) {
		t.Fail()
	}
}

func TestCanBeIncreasing2(t *testing.T) {
	nums := []int{1, 2, 2, 4, 5}
	if !CanBeIncreasing(nums) {
		t.Fail()
	}
}

func TestCanBeIncreasing3(t *testing.T) {
	nums := []int{1, 2, 10, 4, 5}
	if !CanBeIncreasing(nums) {
		t.Fail()
	}
}

func TestCanBeIncreasing4(t *testing.T) {
	nums := []int{2, 3, 1, 2}
	if CanBeIncreasing(nums) {
		t.Fail()
	}
}
