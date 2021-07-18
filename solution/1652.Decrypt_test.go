package solution

import "testing"

func TestDecrypt(t *testing.T) {
	code := []int{5, 7, 1, 4}
	k := 3
	exp := []int{12, 10, 16, 13}
	res := Decrypt(code, k)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestDecrypt2(t *testing.T) {
	code := []int{2, 4, 9, 3}
	k := -2
	exp := []int{12, 5, 6, 13}
	res := Decrypt(code, k)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestDecrypt3(t *testing.T) {
	code := []int{1, 2, 3, 4}
	k := 0
	exp := []int{0, 0, 0, 0}
	res := Decrypt(code, k)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
