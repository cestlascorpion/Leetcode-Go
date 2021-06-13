package solution

import "testing"

func TestSumEvenAfterQueries(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	queries := [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}
	exp := []int{8, 6, 2, 4}
	res := SumEvenAfterQueries(nums, queries)
	if len(res) != len(exp) {
		t.Fail()
	}
	for i := range res {
		if res[i] != exp[i] {
			t.Fail()
		}
	}
}
