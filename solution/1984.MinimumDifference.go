package solution

import "sort"

func MinimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	delta := nums[k-1] - nums[0]
	for i := 0; i <= len(nums)-k; i++ {
		j := i + k - 1
		d := nums[j] - nums[i]
		if d < delta {
			delta = d
		}
	}
	return delta
}
