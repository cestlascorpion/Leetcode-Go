package solution

import "testing"

func DifferenceOfSum(nums []int) int {
	delta := 0
	for i := 0; i < len(nums); i++ {
		n := nums[i]
		delta += n
		for n > 0 {
			// element sum >= digit sum
			delta -= n % 10
			n /= 10
		}
	}
	return delta
}

func TestDiffenceOfSum(t *testing.T) {
	nums := []int{1, 15, 6, 3}
	if DifferenceOfSum(nums) == 9 {
		t.Log("2535. Difference of Sum: case 1 succeeded.")
	} else {
		t.Error("2535. Difference of Sum: case 1 failed.")
	}
	nums = []int{1, 2, 3, 4}
	if DifferenceOfSum(nums) == 0 {
		t.Log("2535. Difference of Sum: case 2 succeeded.")
	} else {
		t.Error("2535. Difference of Sum: case 2 failed.")
	}
}
