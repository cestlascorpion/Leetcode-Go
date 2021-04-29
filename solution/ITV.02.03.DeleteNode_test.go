package solution

import (
	"fmt"
	"testing"
)

func createList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	res := NewListNode(nums[0], nil)
	pre := res
	for i := 1; i < len(nums); i++ {
		cur := NewListNode(nums[i], nil)
		pre.Next = cur
		pre = cur
	}
	return res
}

func TestDeleteNode(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6}
	node := createList(nums)
	DeleteNode(node.Next)
	for node != nil {
		if node.Val == 2 {
			t.Fail()
		}
		fmt.Println(node.Val)
		node = node.Next
	}
}
