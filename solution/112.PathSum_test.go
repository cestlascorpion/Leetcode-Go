package solution

func HasPashSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return targetSum == root.Val
	}
	return HasPashSum(root.Left, targetSum-root.Val) || HasPashSum(root.Right, targetSum-root.Val)
}
