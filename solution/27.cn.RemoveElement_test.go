package solution

import (
	"fmt"
	"testing"
)

func cnRemoveElement(nums []int, val int) int {
	i, j := 0, len(nums)-1
	n := len(nums)
	for i <= j {
		if nums[i] != val {
			i++
		} else {
			if nums[j] != val {
				nums[i] = nums[j]
				j--
				i++
				n--
			} else {
				j--
				n--
			}
		}
	}
	return n
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 3, 3, 3}
	n := cnRemoveElement(nums, 3)
	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}
