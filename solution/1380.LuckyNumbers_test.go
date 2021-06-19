package solution

import "testing"

func TestLuckyNumbers(t *testing.T) {
	matrix := [][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}
	exp := []int{15}
	res := LuckyNumbers(matrix)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
