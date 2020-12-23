package binary_tree

//给定一个二叉树，返回其节点值自底向上的层序遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//
//    3
//   / \
//  9  20
//    /  \
//   15   7
//
//
// 返回其自底向上的层序遍历为：
//
//
//[
//  [15,7],
//  [9,20],
//  [3]
//]

/*
队列:先进先出


*/

func LevelOrderBottom(root *TreeNode) [][]int {
	var levelOrder [][]int
	if root == nil {
		return levelOrder
	}

	var queue []*TreeNode = []*TreeNode{
		root,
	}

	for len(queue) > 0 {
		level := []int{}
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]
			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		levelOrder = append(levelOrder, level)
	}

	var reverse [][]int
	for i := len(levelOrder) - 1; i >= 0; i-- {
		reverse = append(reverse, levelOrder[i])
	}

	return reverse
}
