package solution

func ArraySign(nums []int) int {
	negative := 0
	for i := range nums {
		if nums[i] == 0 {
			return 0
		}
		if nums[i] < 0 {
			negative++
		}
	}
	if negative%2 == 0 {
		return 1
	} else {
		return -1
	}
}
