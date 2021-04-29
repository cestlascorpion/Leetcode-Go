package solution

func MinOperations(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	ops := 0
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > pre {
			pre = nums[i]
		} else if nums[i] == pre {
			ops++
			pre = nums[i] + 1
		} else {
			ops += pre + 1 - nums[i]
			pre = pre + 1
		}
	}
	return ops
}
