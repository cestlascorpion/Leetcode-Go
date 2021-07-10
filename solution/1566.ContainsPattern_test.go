package solution

import "testing"

func TestContainsPattern(t *testing.T) {
	arr := []int{1, 2, 4, 4, 4, 4}
	m := 1
	k := 3
	if !ContainsPattern(arr, m, k) {
		t.Fail()
	}
}

func TestContainsPattern2(t *testing.T) {
	arr := []int{1, 2, 1, 2, 1, 1, 1, 3}
	m := 2
	k := 2
	if !ContainsPattern(arr, m, k) {
		t.Fail()
	}
}

func TestContainsPattern3(t *testing.T) {
	arr := []int{1, 2, 1, 2, 1, 3}
	m := 2
	k := 3
	if ContainsPattern(arr, m, k) {
		t.Fail()
	}
}

func TestContainsPattern4(t *testing.T) {
	arr := []int{2, 2, 2, 2}
	m := 2
	k := 3
	if ContainsPattern(arr, m, k) {
		t.Fail()
	}
}

func TestContainsPattern5(t *testing.T) {
	arr := []int{2,2}
	m := 1
	k := 2
	if !ContainsPattern(arr, m, k) {
		t.Fail()
	}
}

func TestContainsPattern6(t *testing.T) {
	arr := []int{1,2,2,2,1,1,2,2,2}
	m := 1
	k := 3
	if !ContainsPattern(arr, m, k) {
		t.Fail()
	}
}
