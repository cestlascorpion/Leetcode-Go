package solution

func ReversePrint(head *ListNode) []int {
	res := make([]int, 0)
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	max := len(res)
	for i := 0; i < max-i-1; i++ {
		res[i], res[max-i-1] = res[max-i-1], res[i]
	}
	return res
}
