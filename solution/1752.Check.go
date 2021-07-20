package solution

func Check(nums []int) bool {
	h, t := nums[0], nums[len(nums)-1]
	cur := h
	cnt := 0
	for i := range nums {
		if nums[i] < cur {
			cnt++
			if cnt > 1 {
				return false
			}
		}
		cur = nums[i]
	}
	if cnt == 0 {
		return true
	}
	return h >= t
}
