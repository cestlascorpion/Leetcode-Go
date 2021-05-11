package solution

import "testing"

func constructList() *ListNode {
	list := [5]*ListNode{
		&ListNode{1, nil},
		&ListNode{2, nil},
		&ListNode{3, nil},
		&ListNode{4, nil},
		&ListNode{5, nil},
	}
	for i := 1; i < len(list); i++ {
		list[i-1].Next = list[i]
	}
	return list[0]
}

func TestGetKthFromEnd(t *testing.T) {
	list := constructList()
	exp := 4
	if GetKthFromEnd(list, 2).Val != exp {
		t.Fail()
	}
}
