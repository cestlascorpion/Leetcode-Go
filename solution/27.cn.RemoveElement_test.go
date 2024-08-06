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

func cnRemoveElement2(nums []int, val int) int {
	i, j := 0, len(nums)-1
	for i <= j {
		if nums[i] == val {
			for nums[j] == val && j > i {
				j--
			}
			if j == i {
				break
			}
			nums[i], nums[j] = nums[j], nums[i]
		}
		i++
	}
	return i
}

func TestRemoveElement(t *testing.T) {
	nums := []int{3, 3, 3, 3}
	n := cnRemoveElement(nums, 3)
	for i := 0; i < n; i++ {
		fmt.Println(nums[i])
	}
}

func TestRemoveElement2(t *testing.T) {
	nums := []int{3, 3, 3, 3}
	n := cnRemoveElement2(nums, 3)
	if n != 0 {
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		t.FailNow()
	}
	nums = []int{1, 2, 3, 4}
	n = cnRemoveElement2(nums, 5)
	if n != 4 {
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		t.FailNow()
	}
	nums = []int{1, 2, 3, 4}
	n = cnRemoveElement2(nums, 1)
	if n != 3 {
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		t.FailNow()
	}
	nums = []int{1, 2, 3, 4}
	n = cnRemoveElement2(nums, 2)
	if n != 3 {
		t.FailNow()
	}
	nums = []int{1, 2, 3, 4}
	n = cnRemoveElement2(nums, 3)
	if n != 3 {
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		t.FailNow()
	}
	nums = []int{1, 2, 3, 4}
	n = cnRemoveElement2(nums, 4)
	if n != 3 {
		for i := 0; i < n; i++ {
			fmt.Println(nums[i])
		}
		t.FailNow()
	}
}
