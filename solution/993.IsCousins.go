package solution

var xPar, xDep, yPar, yDep int

func isCousinsHelper(node *TreeNode, dep, x, y, pVal int) {
	if node == nil {
		return
	}
	if node.Val == x {
		xPar = pVal
		xDep = dep
	} else if node.Val == y {
		yPar = pVal
		yDep = dep
	} else {
		isCousinsHelper(node.Left, dep+1, x, y, node.Val)
		isCousinsHelper(node.Right, dep+1, x, y, node.Val)
	}
}

func IsCousins(root *TreeNode, x, y int) bool {
	isCousinsHelper(root.Left, 1, x, y, root.Val)
	isCousinsHelper(root.Right, 1, x, y, root.Val)
	return (xPar != yPar) && (xDep == yDep)
}
