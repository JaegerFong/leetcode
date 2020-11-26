package binary_tree

func maxDepth(root *TreeNode) int {
	return DfsDep(root)
}

func DfsDep(root *TreeNode) int{
	if root == nil {
		return 0
	}

	leftDep := DfsDep(root.Left)
	rightDep := DfsDep(root.Right)
	return max(leftDep,rightDep) +  1
}
