package solution

func GetMinDistance(nums []int, target int, start int) int {
	for i := 0; start+i < len(nums) || start-i >= 0; i++ {
		if (start + i) < len(nums) {
			if nums[start+i] == target {
				return i
			}
		}
		if (start - i) >= 0 {
			if nums[start-i] == target {
				return i
			}
		}
	}
	return -1
}
