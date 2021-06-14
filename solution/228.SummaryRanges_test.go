package solution

import (
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	exp := []string{"0->2", "4->5", "7"}
	res := SummaryRanges(nums)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}

func TestSummaryRanges2(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	exp := []string{"0", "2->4", "6", "8->9"}
	res := SummaryRanges(nums)
	if !IsSameStrArray(exp, res) {
		t.Fail()
	}
}
