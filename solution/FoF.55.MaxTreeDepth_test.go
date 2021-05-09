package solution

import "testing"

func constructTree() *TreeNode {
	nodes := []*TreeNode{
		&TreeNode{9, nil, nil},
		&TreeNode{3, nil, nil},
		&TreeNode{20, nil, nil},
		&TreeNode{15, nil, nil},
		&TreeNode{17, nil, nil},
	}
	nodes[1].Left = nodes[0]
	nodes[1].Right = nodes[2]
	nodes[2].Left = nodes[3]
	nodes[2].Right = nodes[4]
	return nodes[1]
}

func TestMaxTreeDepth(t *testing.T) {
	tree := constructTree()
	exp := 3
	if MaxTreeDepth(tree) != exp {
		t.Fail()
	}
}
