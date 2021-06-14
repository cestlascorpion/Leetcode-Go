package solution

import (
	"testing"
)

func constructCousinsTree() *TreeNode {
	node := []*TreeNode{
		&TreeNode{1, nil, nil}, // 0
		&TreeNode{2, nil, nil}, // 1
		&TreeNode{3, nil, nil}, // 2
		&TreeNode{4, nil, nil}, // 3
		&TreeNode{5, nil, nil}, // 4
	}
	node[0].Left = node[1]
	node[0].Right = node[2]
	node[1].Right = node[3]
	node[2].Right = node[4]
	return node[0]
}

func TestIsCousins(t *testing.T) {
	root := constructCousinsTree()
	x, y := 4, 5
	if !IsCousins(root, x, y) {
		t.Fail()
	}
}

func TestIsCousins2(t *testing.T) {
	root := constructCousinsTree()
	x, y := 2, 3
	if IsCousins(root, x, y) {
		t.Fail()
	}
}
