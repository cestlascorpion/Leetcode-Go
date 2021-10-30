package solution

func FindMiddleIndex(nums []int) int {
	sum := 0
	for i := range nums {
		sum += nums[i]
	}
	s := 0
	for i := range nums {
		h := sum - nums[i]
		if h%2 == 0 && s == h/2 {
			return i
		}
		s += nums[i]
	}
	return -1
}
