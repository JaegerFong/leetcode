package binary_tree

import (
	"math"
)

//给出一个完全二叉树，求出该树的节点个数

// 普通的二叉树求节点数
func normalCountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + normalCountNodes(root.Left) + normalCountNodes(root.Right)
}

// 满二叉树的实现
func PerfectCountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh := 0
	l := root
	for l != nil {
		l = l.Left
		lh++
	}

	return int(math.Pow(2, float64(lh))) - 1
}

func CountNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	lh, rh := 1, 1 // 本层算1层
	l, r := root, root

	for l.Left != nil {
		l = l.Left
		lh++
	}

	for r.Right != nil {
		r = r.Right
		rh++
	}

	if lh == rh {
		return int(math.Pow(2, float64(lh))) - 1
	}

	return 1 + CountNodes(root.Left) + CountNodes(root.Right)
}
