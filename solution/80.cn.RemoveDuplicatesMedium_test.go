package solution

import (
	"fmt"
	"testing"
)

func cnRemoveDuplicatesMedium(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	rest := 2
	for cur := 2; cur < len(nums); cur++ {
		if nums[cur] == nums[rest-2] {
			continue
		}
		nums[rest] = nums[cur]
		rest++
	}
	return rest
}

func TestRemoveDuplicatesMedium(t *testing.T) {
	nums := []int{1, 1, 1, 3, 4}
	n := cnRemoveDuplicatesMedium(nums)
	fmt.Println(n, nums)
}
