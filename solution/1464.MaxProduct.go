package solution

func MaxProduct(nums []int) int {
	i, j := nums[0], nums[1]
	if i < j {
		i, j = j, i
	}
	for x := 2; x < len(nums); x++ {
		if nums[x] > i {
			j = i
			i = nums[x]
		} else if nums[x] > j {
			j = nums[x]
		}
	}
	return (i - 1) * (j - 1)
}
