package binary_tree

// 669. 修剪二叉搜索树
func trimBST(root *TreeNode, low int, high int) *TreeNode {
	if root == nil {
		return nil
	}

	// 节点小于low，在需要抛弃左子树
	if root.Val < low {
		root = root.Right
		root = trimBST(root, low, high)
	} else if root.Val > high {
		root = root.Left
		root = trimBST(root, low, high)
	} else {
		// val在区间内，则继续裁剪
		root.Left = trimBST(root.Left, low, high)
		root.Right = trimBST(root.Right, low, high)
	}

	return root
}
