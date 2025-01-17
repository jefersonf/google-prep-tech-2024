package module2

func isBalanced(root *TreeNode) bool {
	ok, _ := checkIsBalanced(root)
	return ok
}

func checkIsBalanced(node *TreeNode) (bool, int) {
	if node == nil {
		return true, 0
	}

	isLeftBalanced, leftHeight := checkIsBalanced(node.Left)
	isRightBalanced, rightHeight := checkIsBalanced(node.Right)

	currHeight := 1 + max(leftHeight, rightHeight)

	if !isLeftBalanced || !isRightBalanced {
		return false, currHeight
	}

	diffHeight := leftHeight - rightHeight
	if diffHeight < 0 {
		diffHeight *= -1
	}

	return diffHeight <= 1, currHeight
}
