package module2

const Inf int64 = (1 << 32)

func isValidBST(root *TreeNode) bool {
	isValid, _, _ := check(root)
	return isValid
}

func check(node *TreeNode) (bool, int64, int64) {
	if node == nil {
		return true, -Inf, Inf
	}

	isLeftvalid, maxLeft, minLeft := check(node.Left)
	isRightValid, maxRight, minRight := check(node.Right)

	var (
		maxVal = max(int64(node.Val), max(maxLeft, maxRight))
		minVal = min(int64(node.Val), min(minLeft, minRight))
	)

	if node.Left != nil && node.Right != nil {
		return isLeftvalid && isRightValid && (maxLeft < int64(node.Val)) && (int64(node.Val) < minRight), maxVal, minVal
	}

	if node.Left != nil {
		return isLeftvalid && isRightValid && (maxLeft < int64(node.Val)), maxVal, minVal
	}

	if node.Right != nil {
		return isLeftvalid && isRightValid && (int64(node.Val) < minRight), maxVal, minVal
	}

	return true, maxVal, minVal
}
