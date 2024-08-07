package solution

import (
	"testing"
)

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(height(root.Left), height(root.Right)) + 1
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if abs(height(root.Left)-height(root.Right)) > 1 {
		return false
	}
	return IsBalanced(root.Left) && IsBalanced(root.Right)
}

func heightAndIsBalanced(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	leftHeight, leftIsBalanced := heightAndIsBalanced(root.Left)
	rightHeight, rightIsBalanced := heightAndIsBalanced(root.Right)
	if !leftIsBalanced || !rightIsBalanced || abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}
	return max(leftHeight, rightHeight) + 1, true
}

func IsBalancedFast(root *TreeNode) bool {
	_, isBalanced := heightAndIsBalanced(root)
	return isBalanced
}

func TestIsBalanced(t *testing.T) {
	tree := &TreeNode{3, &TreeNode{9, nil, nil}, &TreeNode{20, &TreeNode{15, nil, nil}, &TreeNode{7, nil, nil}}}
	t.Log(IsBalanced(tree))
}
