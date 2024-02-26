package solution

import (
	"fmt"
	"testing"
)

func cnTwoSum(nums []int, target int) []int {
	m := make(map[int]int)
	m[nums[0]] = 0
	for i := 1; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok {
			return []int{v, i}
		} else {
			m[nums[i]] = i
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	nums := []int{3, 2, 4}
	fmt.Println(cnTwoSum(nums, 6))
}
