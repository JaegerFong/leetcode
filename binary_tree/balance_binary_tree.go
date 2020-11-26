package binary_tree

import "math"

//输入一棵二叉树的根节点，求该树的深度。从根节点到叶节点依次经过的节点（含根、叶节点）形成树的一条路径，最长路径的长度为树的深度。
func isBalanced(root *TreeNode) bool {
	if postorder(root) >= 0 {
		return true
	}
	return false
}

func postorder(root *TreeNode) int{
	if root == nil {
		return 0
	}

	leftDep := postorder(root.Left)
	rightDep := postorder(root.Right)

	if leftDep >=0 && rightDep >= 0 && math.Abs( float64(leftDep - rightDep)) <=1 {
		return max(leftDep,rightDep) + 1
	}else{
		return -1
	}

}

func max(left,right int)int{
	if left > right{
		return left
	}

	return right
}