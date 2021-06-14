package solution

import "testing"

func constructSumRootToLeafTree() *TreeNode {
	nodes := []*TreeNode{
		&TreeNode{1, nil, nil}, // 0
		&TreeNode{0, nil, nil}, // 1
		&TreeNode{0, nil, nil}, // 2
		&TreeNode{1, nil, nil}, // 3
		&TreeNode{1, nil, nil}, // 4
		&TreeNode{0, nil, nil}, // 5
		&TreeNode{1, nil, nil}, // 6
	}
	nodes[0].Left = nodes[1]
	nodes[0].Right = nodes[4]
	nodes[1].Left = nodes[2]
	nodes[1].Right = nodes[3]
	nodes[4].Left = nodes[5]
	nodes[4].Right = nodes[6]

	return nodes[0]
}

func TestSumRootToLeaf(t *testing.T) {
	root := constructSumRootToLeafTree()
	exp := 22
	if SumRootToLeaf(root) != exp {
		t.Fail()
	}
}

func TestSumRootToLeaf2(t *testing.T) {
	root := &TreeNode{1, nil, &TreeNode{1, nil, nil}}
	exp := 3
	if SumRootToLeaf(root) != exp {
		t.Fail()
	}
}
