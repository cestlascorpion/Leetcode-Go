package solution

func KLengthApart(nums []int, k int) bool {
	pos := -1
	for i := range nums {
		if nums[i] == 1 {
			if pos != -1 {
				if i-pos-1 < k {
					return false
				}
			}
			pos = i
		}
	}
	return true
}
