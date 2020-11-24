package binary_tree

// 翻转二叉树
func InvertTree(root *TreeNode) *TreeNode {

	if root == nil {
		return nil
	}

	InvertTree(root.Left)
	InvertTree(root.Right)

	// 翻转
	root.Left,root.Right = root.Right,root.Left

	return root
}