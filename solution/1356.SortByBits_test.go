package solution

import (
	"fmt"
	"testing"
)

func TestCountBits(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	exp := []int{0, 1, 1, 2, 1, 2, 2, 3, 1}
	for i := range arr {
		if exp[i] != CountBits(arr[i]) {
			t.Fail()
		}
	}
}

func TestSortByBits(t *testing.T) {
	arr := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	exp := []int{0, 1, 2, 4, 8, 3, 5, 6, 7}
	res := SortByBits(arr)
	if !IsSameArray(exp, res) {
		fmt.Println(exp)
		fmt.Println(res)
		t.Fail()
	}
}
