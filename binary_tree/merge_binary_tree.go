package binary_tree
//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点
//。

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	node := new(TreeNode)
	// 其中一个为nil
	if t1 == nil && t2 != nil {
		node.Val = t2.Val
		node.Left = t2.Left
		node.Right = t2.Right
		return node
	}

	if t2 == nil && t1 != nil {
		node.Val = t1.Val
		node.Left = t1.Left
		node.Right = t1.Right
		return node
	}


	node.Val = t1.Val + t2.Val
	node.Left = mergeTrees(t1.Left,t2.Left)
	node.Right = mergeTrees(t1.Right,t2.Right)

	return node
}