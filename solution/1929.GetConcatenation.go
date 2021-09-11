package solution

func GetConcatenation(nums []int) []int {
	ans := make([]int, len(nums)*2)
	for i := range nums {
		ans[i] = nums[i]
	}
	for i := range nums {
		ans[len(nums)+i] = nums[i]
	}
	return ans
}
