package solution

func CountKDifference(nums []int, k int) int {
	ans := 0
	m := make(map[int]int)
	for i := range nums {
		ans += m[nums[i]+k] + m[nums[i]-k]
		m[nums[i]]++
	}
	return ans
}

func CountKDifference2(nums []int, k int) int {
	count := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			delta := nums[i] - nums[j]
			if delta == k || delta == -k {
				count++
			}
		}
	}
	return count
}
