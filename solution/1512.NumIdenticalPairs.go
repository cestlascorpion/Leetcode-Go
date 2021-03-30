package solution

func NumIdenticalPairs(nums []int) int {
	ret := 0
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] {
				ret++
			}
		}
	}
	return ret
}

func NumIdenticalPairs2(nums []int) int {
	m := make(map[int]int, 0)
	for _, v := range nums {
		m[v]++
	}
	ret := 0
	for _, v := range m {
		ret += v * (v - 1) / 2
	}
	return ret
}
