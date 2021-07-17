package solution

func GetMaximumGenerated(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	res := 1
	nums := make([]int, n+1)
	nums[0] = 0
	nums[1] = 1
	for i := 2; i < len(nums); i++ {
		if (i % 2) == 0 {
			nums[i] = nums[i/2]

		} else {
			j := (i - 1) / 2
			nums[i] = nums[j] + nums[j+1]
		}
		if nums[i] > res {
			res = nums[i]
		}
	}
	return res
}
