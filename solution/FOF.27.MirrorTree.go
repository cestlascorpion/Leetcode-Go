package solution

func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	left := MirrorTree(root.Right)
	right := MirrorTree(root.Left)
	return &TreeNode{root.Val, left, right}
}

func MirrorTreeMutable(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Right, root.Left = root.Left, root.Right
	MirrorTreeMutable(root.Right)
	MirrorTreeMutable(root.Left)
	return root
}