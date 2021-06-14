package solution

func sumRootToLeafHelper(node *TreeNode, current int, sum *int) {
	current = (current << 1) + node.Val
	if node.Left == nil && node.Right == nil {
		*sum += current
		return
	}
	if node.Left != nil {
		sumRootToLeafHelper(node.Left, current, sum)
	}
	if node.Right != nil {
		sumRootToLeafHelper(node.Right, current, sum)
	}
}

func SumRootToLeaf(root *TreeNode) int {
	sum := 0
	if root == nil {
		return sum
	}
	sumRootToLeafHelper(root, 0, &sum)
	return sum
}
