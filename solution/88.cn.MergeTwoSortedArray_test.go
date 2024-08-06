package solution

import (
	"fmt"
	"sort"
	"testing"
)

func cnMerge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func cnMergeFast(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	i := m - 1
	j := n - 1
	x := m + n - 1
	for {
		if i < 0 {
			for j >= 0 {
				nums1[x] = nums2[j]
				x--
				j--
			}
			break
		}
		if j < 0 {
			for i >= 0 {
				nums1[x] = nums1[i]
				x--
				i--
			}
			break
		}
		if nums1[i] >= nums2[j] {
			nums1[x] = nums1[i]
			x--
			i--
		} else {
			nums1[x] = nums2[j]
			x--
			j--
		}
	}
}

func cnMergeFast2(nums1 []int, m int, nums2 []int, n int) {
	if len(nums1) != m+n {
		return
	}
	i, j, x := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] >= nums2[j] {
			nums1[x] = nums1[i]
			i--
			x--
		} else {
			nums1[x] = nums2[j]
			j--
			x--
		}
	}
	for i >= 0 {
		nums1[x] = nums1[i]
		i--
		x--
	}
	for j >= 0 {
		nums1[x] = nums2[j]
		j--
		x--
	}
}

func TestMerge(t *testing.T) {
	nums1 := make([]int, 8)
	nums1[0] = 1
	nums1[1] = 3
	nums1[2] = 4
	nums1[3] = 7
	nums1[4] = 8
	nums2 := make([]int, 3)
	nums2[0] = 2
	nums2[1] = 3
	nums2[2] = 5
	cnMerge(nums1, 5, nums2, 3)
	fmt.Println(nums1)

}

func TestMergeFast(t *testing.T) {
	nums1 := make([]int, 8)
	nums1[0] = 1
	nums1[1] = 3
	nums1[2] = 4
	nums1[3] = 7
	nums1[4] = 8
	nums2 := make([]int, 3)
	nums2[0] = 2
	nums2[1] = 3
	nums2[2] = 5
	cnMergeFast(nums1, 5, nums2, 3)
	fmt.Println(nums1)
}

func TestMergeFast2(t *testing.T) {
	nums1 := make([]int, 8)
	nums1[0] = 1
	nums1[1] = 3
	nums1[2] = 4
	nums1[3] = 7
	nums1[4] = 8
	nums2 := make([]int, 3)
	nums2[0] = 2
	nums2[1] = 3
	nums2[2] = 5
	cnMergeFast2(nums1, 5, nums2, 3)
	fmt.Println(nums1)
}
