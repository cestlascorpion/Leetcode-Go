package solution

import (
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	exp := []int{2, 0, 3}
	res := KWeakestRows(mat, 3)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestKWeakestRows2(t *testing.T) {
	mat := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1}}
	exp := []int{2, 0, 3}
	res := KWeakestRows2(mat, 3)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}

func TestKWeakestRows3(t *testing.T) {
	mat := [][]int{
		{1, 1, 0},
		{1, 1, 0},
		{1, 1, 1},
		{1, 1, 1},
		{0, 0, 0},
		{1, 1, 1},
		{1, 0, 0}}
	exp := []int{4, 6, 0, 1, 2, 3}
	res := KWeakestRows2(mat, 6)
	if !IsSameArray(exp, res) {
		t.Fail()
	}
}
