package solution

import (
	"math"
	"testing"
)

func MinimumCost(nums []int) int {
	// just find the two smallest numbers in the array[1:]
	x, y := math.MaxInt, math.MaxInt
	for i := 1; i < len(nums); i++ {
		if nums[i] < x && nums[i] < y {
			if x > y {
				x = nums[i]
			} else {
				y = nums[i]
			}
		} else if nums[i] < x && nums[i] >= y {
			x = nums[i]
		} else if nums[i] >= x && nums[i] < y {
			y = nums[i]
		}
	}
	return x + y + nums[0]
}

func TestMinimumCost(t *testing.T) {
	nums := []int{1, 2, 3, 12}
	if MinimumCost(nums) == 6 {
		t.Log("3010. Divide an Array Into Subarrays With Minimum Cost I: case 1 succeeded.")
	} else {
		t.Error("3010. Divide an Array Into Subarrays With Minimum Cost I: case 1 failed.")
	}
	nums = []int{5, 3, 4}
	if MinimumCost(nums) == 12 {
		t.Log("3010. Divide an Array Into Subarrays With Minimum Cost I: case 2 succeeded.")
	} else {
		t.Error("3010. Divide an Array Into Subarrays With Minimum Cost I: case 2 failed.")
	}
	nums = []int{10, 3, 1, 1}
	if MinimumCost(nums) == 12 {
		t.Log("3010. Divide an Array Into Subarrays With Minimum Cost I: case 3 succeeded.")
	} else {
		t.Error("3010. Divide an Array Into Subarrays With Minimum Cost I: case 3 failed.")
	}
}
