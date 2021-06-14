package solution

import "testing"

func TestAddToArrayForm(t *testing.T) {
	A := []int{1, 2, 0, 0}
	K := 34
	exp := []int{1, 2, 3, 4}
	res := AddToArrayForm(A, K)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestAddToArrayForm2(t *testing.T) {
	A := []int{2, 7, 4}
	K := 181
	exp := []int{4, 5, 5}
	res := AddToArrayForm(A, K)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestAddToArrayForm3(t *testing.T) {
	A := []int{2, 1, 5}
	K := 806
	exp := []int{1, 0, 2, 1}
	res := AddToArrayForm(A, K)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestAddToArrayForm4(t *testing.T) {
	A := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 0, 6, 3}
	K := 516
	exp := []int{1, 2, 6, 3, 0, 7, 1, 7, 1, 9, 7, 5, 6, 6, 4, 4, 0, 5, 7, 9}
	res := AddToArrayForm(A, K)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
