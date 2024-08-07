package solution

import "testing"

func IsSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

func TestIsSameTree(t *testing.T) {
	tree1 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	tree2 := &TreeNode{1, &TreeNode{2, nil, nil}, &TreeNode{3, nil, nil}}
	t.Log(IsSameTree(tree1, tree2))
}