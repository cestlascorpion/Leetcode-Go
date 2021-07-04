package solution

import "sort"

func CanMakeArithmeticProgression(arr []int) bool {
	if len(arr) < 3 {
		return true
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	delta := arr[1] - arr[0]
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != delta {
			return false
		}
	}
	return true
}
