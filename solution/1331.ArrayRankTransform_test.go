package solution

import "testing"

func TestArrayRankTransform(t *testing.T) {
	arr := []int{4, 1, 2, 3}
	exp := []int{4, 1, 2, 3}
	res := ArrayRankTransform(arr)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestArrayRankTransform2(t *testing.T) {
	arr := []int{100, 100, 100}
	exp := []int{1, 1, 1}
	res := ArrayRankTransform(arr)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestArrayRankTransform3(t *testing.T) {
	arr := []int{37,12,28,9,100,56,80,5,12}
	exp := []int{5,3,4,2,8,6,7,1,3}
	res := ArrayRankTransform(arr)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
