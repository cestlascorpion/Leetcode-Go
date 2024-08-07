package solution

import "testing"

func DeleteDuplicates(head *ListNode) *ListNode {
	pos := head
	for pos != nil {
		for pos.Next != nil && pos.Val == pos.Next.Val {
			pos.Next = pos.Next.Next
		}
		pos = pos.Next
	}
	return head
}

func TestDeleteDuplicates(t *testing.T) {
	array:= []ListNode{
		{1, nil},
		{2, nil},
		{3, nil},
		{3, nil},
		{4, nil},
		{4, nil},
		{5, nil},
		{5, nil},
	}
	for i := 0; i < len(array) - 1; i++ {
		array[i].Next = &array[i + 1]
	}
	head := DeleteDuplicates(&array[0])
	for head != nil {
		t.Log(head.Val)
		head = head.Next
	}
}