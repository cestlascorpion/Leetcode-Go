package solution

func MinStartValue(nums []int) int {
	min, pre := nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		pre = pre + nums[i]
		if pre < min {
			min = pre
		}
	}
	res := 1 - min
	if res <= 0 {
		res = 1
	}
	return res
}
