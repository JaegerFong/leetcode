package binary_tree
func PreorderTraversal(root *TreeNode) (vals []int) {
	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node == nil {
			return
		}

		vals = append(vals,node.Val)
		preOrder(node.Left)
		preOrder(node.Right)
	}
	preOrder(root)
	return vals
}