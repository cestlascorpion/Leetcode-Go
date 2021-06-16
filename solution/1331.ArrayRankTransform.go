package solution

import "sort"

func ArrayRankTransform(arr []int) []int {
	res := make([]int, len(arr))
	if len(arr) == 0 {
		return res
	}
	max := arr[0]
	min := arr[0]
	for i := range arr {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}
	bucket := make([]int, max-min+1)
	for i := range arr {
		bucket[arr[i]-min]++
	}
	rank := 1
	for i := range bucket {
		if bucket[i] != 0 {
			bucket[i] = rank
			rank++
		}
	}
	for i := range arr {
		res[i] = bucket[arr[i]-min]
	}
	return res
}

func ArrayRankTransform2(arr []int) []int {
	res := make([]int, len(arr))
	n := len(arr)
	if n == 0 {
		return res
	}
	idx := make([]int, n)
	for i := range idx {
		idx[i] = i
	}
	sort.Slice(idx, func(n1, n2 int) bool {
		return arr[idx[n1]] < arr[idx[n2]]
	})
	res[idx[0]] = 1
	for i := 1; i < n; i++ {
		if arr[idx[i]] == arr[idx[i-1]] {
			res[idx[i]] = res[idx[i-1]]
		} else {
			res[idx[i]] = res[idx[i-1]] + 1
		}
	}
	return res
}
