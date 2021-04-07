package solution

func CreateTargetArray(nums []int, index []int) []int {
	res := make([]int, len(nums), len(nums))
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		idx := index[i]
		for cur := i; cur > idx; cur-- {
			res[cur] = res[cur-1]
		}
		res[idx] = num
	}
	return res
}
