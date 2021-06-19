package solution

import (
	"sort"
)

func CountBits(num int) int {
	cnt := 0
	for num != 0 {
		if (num & 1) == 1 {
			cnt++
		}
		num = num >> 1
	}
	return cnt
}

func SortByBits(arr []int) []int {
	sort.Slice(arr, func(i, j int) bool {
		b1, b2 := CountBits(arr[i]), CountBits(arr[j])
		if b1 == b2 {
			return arr[i] < arr[j]
		} else {
			return b1 < b2
		}
	})
	return arr
}
