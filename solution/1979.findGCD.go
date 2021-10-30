package solution

func FindGCD(nums []int) int {
	min, max := nums[0], nums[0]
	for i := range nums {
		if min > nums[i] {
			min = nums[i]
		}
		if max < nums[i] {
			max = nums[i]
		}
	}
	return gcd(min, max)
}

func gcd(a, b int) int {
	for a != b {
		if a > b {
			a = a - b
		} else if a < b {
			b = b - a
		}
	}
	return a
}
