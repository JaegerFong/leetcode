package binary_tree

import (
	"math"
)

//给出一个完全二叉树，求出该树的节点个数
// 参考链接：https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247485057&idx=1&sn=45a3b89a4efef236cb662d5505d7ce36&chksm=9bd7f889aca0719f4915de681f983355e187151030991ab1944494ffe4b73e484068b85eb01e&scene=21#wechat_redirect
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
