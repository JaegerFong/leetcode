package binary_tree

var sum int

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	sum = 0
	inorder_convert_bst(root)
	return root
}

func inorder_convert_bst(root *TreeNode) {
	if root == nil {
		return
	}
	inorder_convert_bst(root.Right)
	sum += root.Val
	root.Val = sum
	inorder_convert_bst(root.Left)
}
