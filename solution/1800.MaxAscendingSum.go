package solution

func MaxAscendingSum(nums []int) int {
	res := nums[0]
	sum := nums[0]
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > prev {
			sum += nums[i]
		} else {
			sum = nums[i]
		}
		if sum > res {
			res = sum
		}
		prev = nums[i]
	}
	return res
}
