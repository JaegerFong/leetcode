package dp

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func robTree(root *TreeNode) int {
	return robTree(root)
}

func robTreeDp(root *TreeNode) int {
	tmpl := 0
	tmpr := 0
	if root.Left != nil {
		tmpl = robTreeDp(root.Left.Left) + robTreeDp(root.Left.Right)
	}

	if root.Right != nil {
		tmpr = robTreeDp(root.Right.Left) + robTreeDp(root.Right.Right)
	}

	doit := root.Val + tmpl + tmpr
	notdoit := robTreeDp(root.Left) + robTreeDp(root.Right)
	res := max(doit, notdoit)
	return res
}
