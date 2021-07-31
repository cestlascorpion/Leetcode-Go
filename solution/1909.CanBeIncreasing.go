package solution

func check(idx int, nums []int) bool {
	for i := 1; i < len(nums)-1; i++ {
		prev := i - 1
		if prev >= idx {
			prev++
		}
		cur := i
		if cur >= idx {
			cur++
		}
		if nums[cur] <= nums[prev] {
			return false
		}
	}
	return true
}

func CanBeIncreasing(nums []int) bool {
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] <= nums[i-1] {
			return check(i, nums) || check(i-1, nums)
		}
	}
	return true
}
