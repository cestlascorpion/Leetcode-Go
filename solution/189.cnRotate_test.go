package solution

import (
	"fmt"
	"testing"
)

func rotate(nums []int, k int) {
	if k%len(nums) == 0 {
		return
	}
	tmp := make([]int, len(nums))
	for i := range nums {
		tmp[(i+k)%len(nums)] = nums[i]
	}
	copy(nums, tmp)
}

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	rotate(nums, 3)
	fmt.Print(nums)
}
