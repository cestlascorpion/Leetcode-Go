package solution

import (
	"testing"
)

func TestArraySign(t *testing.T) {
	nums := []int{-1, -2, -3, -4, 3, 2, 1}
	exp := 1
	if ArraySign(nums) != exp {
		t.Fail()
	}
}

func TestArraySign2(t *testing.T) {
	nums := []int{1, 5, 0, 2, -3}
	exp := 0
	if ArraySign(nums) != exp {
		t.Fail()
	}
}

func TestArraySign3(t *testing.T) {
	nums := []int{-1, 1, -1, 1, -13}
	exp := -1
	if ArraySign(nums) != exp {
		t.Fail()
	}
}
