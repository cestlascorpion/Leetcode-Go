package solution

func GetKthFromEnd(head *ListNode, k int) *ListNode {
	f, s := head, head
	for k > 0 {
		f = f.Next
		k--
	}
	for f != nil {
		f = f.Next
		s = s.Next
	}
	return s
}
