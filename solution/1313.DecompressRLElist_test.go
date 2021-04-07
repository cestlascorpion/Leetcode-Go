package solution

import "testing"

func TestDecompressRLElist(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	res := DecompressRLElist(nums)
	exp := []int{2, 4, 4, 4}
	if len(res) != len(exp) {
		t.Fail()
	}
	for i, v := range exp {
		if res[i] != v {
			t.Fail()
		}
	}
}
