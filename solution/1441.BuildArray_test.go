package solution

import "testing"

func TestBuildArray(t *testing.T) {
	target := []int{1, 3}
	n := 3
	exp := []string{Push, Push, Pop, Push}
	res := BuildArray(target, n)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}

func TestBuildArray2(t *testing.T) {
	target := []int{1, 2, 3}
	n := 3
	exp := []string{Push, Push, Push}
	res := BuildArray(target, n)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}

func TestBuildArray3(t *testing.T) {
	target := []int{2, 3, 4}
	n := 4
	exp := []string{Push, Pop, Push, Push, Push}
	res := BuildArray(target, n)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}

func TestBuildArray4(t *testing.T) {
	target := []int{1, 2}
	n := 4
	exp := []string{Push, Push}
	res := BuildArray(target, n)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}
