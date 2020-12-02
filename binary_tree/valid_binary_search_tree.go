package binary_tree

// 判断是否有效BST
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	return validBST(root, 0, 0)
}

func validBST(root *TreeNode, min, max int) bool {
	if root == nil {
		return true
	}

	if min != 0 && root.Val <= min {
		return false
	}

	if max != 0 && root.Val >= max {
		return false
	}

	return validBST(root.Left, 0, root.Val) && validBST(root.Right, root.Val, 0)
}
