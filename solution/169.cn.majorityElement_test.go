package solution

import (
	"fmt"
	"testing"
)

func cnMajorityElement(nums []int) int {
	ret := nums[0]
	m := make(map[int]int)
	for i, v := range nums {
		m[v]++
		if i > len(nums)/2 && m[v] > len(nums)/2 {
			ret = v
			break
		}
	}
	return ret
}

func cnMajorityElementFast(nums []int) int {
	ret := nums[0]
	cnt := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == ret {
			cnt++
		} else {
			cnt--
		}
		if cnt == -1 {
			cnt = 1
			ret = nums[i]
		}
	}
	return ret
}

func TestMajorityElementFast(t *testing.T) {
	nums := []int{1, 1, 2}
	fmt.Println(cnMajorityElement(nums))
	fmt.Println(cnMajorityElementFast(nums))
}
