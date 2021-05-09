package solution

func ldr(node *TreeNode, ch chan<- int) {
	if node.Left != nil {
		ldr(node.Left, ch)
	}
	ch <- node.Val
	if node.Right != nil {
		ldr(node.Right, ch)
	}
}

func LDR(node *TreeNode, ch chan<- int) {
	defer close(ch)
	ldr(node, ch)
}
