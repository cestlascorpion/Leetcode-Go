package solution

import "testing"

func TestCheckIfExist(t *testing.T) {
	arr := []int{10, 2, 5, 3}
	if !CheckIfExist(arr) {
		t.Fail()
	}
}

func TestCheckIfExist2(t *testing.T) {
	arr := []int{3, 1, 7, 11}
	if CheckIfExist(arr) {
		t.Fail()
	}
}

func TestCheckIfExist3(t *testing.T) {
	arr := []int{4, -7, 11, 4, 18}
	if CheckIfExist(arr) {
		t.Fail()
	}
}
