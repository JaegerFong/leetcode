package binary_tree

import "math"

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