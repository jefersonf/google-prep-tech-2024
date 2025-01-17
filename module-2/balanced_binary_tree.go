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

	d := leftHeight - rightHeight
	if d < 0 {
		d *= -1
	}

	return d <= 1, currHeight
}
