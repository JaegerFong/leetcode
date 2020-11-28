package binary_tree

func isUnivalTree(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return check(root,root.Val)
}

func check(root *TreeNode , value int) bool{
	if root == nil{
		return true
	}

	if (root.Left != nil && root.Left.Val != value) || (root.Right != nil && root.Right.Val !=  value) {
		return false
	}

	return check(root.Left,value) && check(root.Right,value)
}