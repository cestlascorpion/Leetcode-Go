package solution

import "testing"

/*
      3
   9    20
      15  17
*/
func constructMirrorTree() *TreeNode {
	nodes := []*TreeNode{
		&TreeNode{9, nil, nil},
		&TreeNode{3, nil, nil},
		&TreeNode{16, nil, nil},
		&TreeNode{15, nil, nil},
		&TreeNode{17, nil, nil},
	}
	nodes[1].Left = nodes[0]
	nodes[1].Right = nodes[2]
	nodes[2].Left = nodes[3]
	nodes[2].Right = nodes[4]
	return nodes[1]
}

func TestMirrorTree(t *testing.T) {
	root := constructTree()
	LDR2(root)
	mirror := MirrorTree(root)
	LDR2(mirror)
}

func TestMirrorTreeMutable(t *testing.T) {
	root := constructTree()
	LDR2(root)
	mirror := MirrorTreeMutable(root)
	LDR2(mirror)
}
