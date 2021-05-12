package solution

func maxTreeDepth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := maxTreeDepth(node.Left)
	right := maxTreeDepth(node.Right)
	if left > right {
		return left + 1
	} else {
		return right + 1
	}
}

func MaxTreeDepth(root *TreeNode) int {
	return maxTreeDepth(root)
}
