package solution

import (
	"sort"
)

func KWeakestRows(mat [][]int, k int) []int {
	cnt := make([]int, len(mat))
	for i := 0; i < len(mat); i++ {
		for j := 0; j < len(mat[i]); j++ {
			if mat[i][j] == 1 {
				cnt[i]++
			} else {
				break
			}
		}
	}
	idx := make([]int, len(mat))
	for i := range idx {
		idx[i] = i
	}
	sort.SliceStable(idx, func(i, j int) bool {
		return cnt[idx[i]] < cnt[idx[j]]
	})
	return idx[:k]
}

func KWeakestRows2(mat [][]int, k int) []int {
	res := make([]int, 0)
	for col := 0; col <= len(mat[0]); col++ {
		for row := 0; row < len(mat); row++ {
			if col == len(mat[0]) && mat[row][col-1] != 0 {
				res = append(res, row)
				if len(res) == k {
					return res
				}
			}
			if col <len(mat[0]) && mat[row][col] == 0 && (col == 0 || (col > 0 && mat[row][col-1] != 0)) {
				res = append(res, row)
				if len(res) == k {
					return res
				}
			}
		}
	}
	return res
}
