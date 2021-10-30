package solution

import (
	"sort"
)

func CountQuadruplets(nums []int) int {
	count := 0
	set := map[int]int{}
	for i := 1; i < len(nums)-2; i++ {
		for j := 0; j < i; j++ {
			set[nums[i]+nums[j]]++
		}
		for j := i + 2; j < len(nums); j++ {
			if set[nums[j]-nums[i+1]] > 0 {
				count += set[nums[j]-nums[i+1]]
			}
		}
	}
	return count
}

func CountQuadruplets2(nums []int) int {
	count := 0
	m := make(map[int][]int)
	for i := range nums {
		if val, ok := m[nums[i]]; !ok {
			m[nums[i]] = []int{i}
		} else {
			m[nums[i]] = append(val, i)
		}
	}
	for a := 0; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			for c := b + 1; c < len(nums)-1; c++ {
				sum := nums[a] + nums[b] + nums[c]
				if val, ok := m[sum]; ok {
					idx := sort.Search(len(val), func(i int) bool {
						return val[i] > c
					})
					count += len(val) - idx

				}
			}
		}
	}
	return count
}

func CountQuadruplets3(nums []int) int {
	count := 0
	for a := 0; a < len(nums)-3; a++ {
		for b := a + 1; b < len(nums)-2; b++ {
			for c := b + 1; c < len(nums)-1; c++ {
				for d := c + 1; d < len(nums); d++ {
					if nums[a]+nums[b]+nums[c] == nums[d] {
						count++
					}
				}
			}
		}
	}
	return count
}
