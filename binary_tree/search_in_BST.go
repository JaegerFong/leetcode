package binary_tree

// BST中搜索值,根据二叉搜索树特性，无需遍历整颗树
func isInBST(root *TreeNode, target int) bool {
	if root == nil {
		return false
	}

	if root.Val == target {
		return true
	}

	if target > root.Val {
		return isInBST(root.Right, target)
	}

	if target < root.Val {
		return isInBST(root.Left, target)
	}

	return false
}
