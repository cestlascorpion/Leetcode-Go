package solution

import "testing"

func TestDecompressRLElist(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := DecompressRLElist(nums)
	exp := []int{2, 4, 4, 4}
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
