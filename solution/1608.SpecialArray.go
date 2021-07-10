package solution

func SpecialArray(nums []int) int {
	max := len(nums)
	count := make(map[int]int)
	for i := range nums {
		count[nums[i]]++
		if max < nums[i] {
			max = nums[i]
		}
	}
	sum := 0
	for x := 0; x <= max; x++ {
		if v, ok := count[x]; ok {
			if len(nums)-sum == x {
				return x
			}
			sum += v
		} else {
			if len(nums)-sum == x {
				return x
			}
		}
	}
	return -1
}