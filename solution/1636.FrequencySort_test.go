package solution

import "testing"

func TestFrequencySortNums(t *testing.T) {
	nums := []int{1, 1, 2, 2, 2, 3}
	exp := []int{3, 1, 1, 2, 2, 2}
	res := FrequencySortNums(nums)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestFrequencySortNums2(t *testing.T) {
	nums := []int{2,3,1,3,2}
	exp := []int{1,3,3,2,2}
	res := FrequencySortNums(nums)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
