package solution

import (
	"sort"
)

func FrequencySortNums(nums []int) []int {
	dict := make(map[int]int, 0)
	for i := range nums {
		dict[nums[i]]++
	}
	s := make([][2]int, 0, len(dict))
	for k, v := range dict {
		s = append(s, [2]int{k, v})
	}
	sort.Slice(s, func(i, j int) bool {
		if s[i][1] == s[j][1] {
			return s[i][0] > s[j][0]
		}
		return s[i][1] < s[j][1]
	})
	res := make([]int, 0, len(nums))
	for i := range s {
		for j := 0; j < s[i][1]; j++ {
			res = append(res, s[i][0])
		}
	}
	return res
}

func FrequencySortNums2(nums []int) []int {
	maps := make(map[int]int, 0)
	for _, num := range nums {
		maps[num]++
	}
	sort.Slice(nums, func(i, j int) bool {
		if maps[nums[i]] == maps[nums[j]] {
			return nums[i] > nums[j]
		}
		return maps[nums[i]] < maps[nums[j]]
	})
	return nums
}
