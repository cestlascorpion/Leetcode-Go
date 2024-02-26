package solution

import (
	"fmt"
	"testing"
)

func cnRemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dup := make([]int, 0, len(nums))
	dup = append(dup, nums[0])
	j := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > dup[j] {
			dup = append(dup, nums[i])
			j++
		}
	}
	copy(nums, dup)
	return len(dup)
}

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	n := cnRemoveDuplicates(nums)
	fmt.Println(n, nums)
}

func cnRemoveDuplicatesFast(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	rest := 1
	for cur := 1; cur < len(nums); cur++ {
		if nums[cur] == nums[rest-1] {
			continue
		}
		nums[rest] = nums[cur]
		rest++
	}
	return rest
}

func TestRemoveDuplicatesFast(t *testing.T) {
	nums := []int{1, 1, 1, 1, 1}
	n := cnRemoveDuplicatesFast(nums)
	fmt.Println(n, nums)
}
