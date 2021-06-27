package solution

import "sort"

func MinSubsequence(nums []int) []int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	res := make([]int, 0)
	subSum := 0
	for i := range nums {
		subSum += nums[i]
		res = append(res, nums[i])
		if subSum > sum/2 {
			break
		}
	}
	return res
}
