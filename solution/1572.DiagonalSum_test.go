package solution

import (
	"fmt"
	"testing"
)

func TestDiagonalSum(t *testing.T) {
	mat := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	res := DiagonalSum(mat)
	exp := 25
	if exp != res {
		fmt.Println(res, exp)
		t.Fail()
	}
}

func TestDiagonalSum2(t *testing.T) {
	mat := [][]int{{1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}, {1, 1, 1, 1}}
	res := DiagonalSum(mat)
	exp := 8
	if exp != res {
		fmt.Println(res, exp)
		t.Fail()
	}
}
