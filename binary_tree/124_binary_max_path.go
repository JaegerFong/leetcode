package binary_tree

import "math"

//给定一个非空二叉树，返回其最大路径和。
// 本题中，路径被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。该路径至少包含一个节点，且不一定经过根节点。
//
//
//
// 示例 1：
//
// 输入：[1,2,3]
//
//       1
//      / \
//     2   3
//
//输出：6
//

// 后续遍历

func MaxPathSum(root *TreeNode) int {
	curMax := 0
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if root == nil {
			return 0
		}

		l := max(0, maxGain(root.Left))
		r := max(0, maxGain(root.Right))
		curMax = max(curMax, l+r+root.Val)
		return max(l, r) + root.Val
	}
	maxGain(root)
	return curMax
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
