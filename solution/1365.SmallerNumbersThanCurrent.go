package solution

func SmallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, len(nums))
	count := [101]int{}
	for i := range nums {
		count[nums[i]]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for i := range nums {
		if nums[i] != 0 {
			res[i] = count[nums[i]-1]
		}
	}
	return res
}
