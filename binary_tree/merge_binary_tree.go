package binary_tree
//给定两个二叉树，想象当你将它们中的一个覆盖到另一个上时，两个二叉树的一些节点便会重叠。
//
// 你需要将他们合并为一个新的二叉树。合并的规则是如果两个节点重叠，那么将他们的值相加作为节点合并后的新值，否则不为 NULL 的节点将直接作为新二叉树的节点
//。

func mergeTrees(t1 *TreeNode, t2 *TreeNode) *TreeNode {
	if t1 == nil && t2 == nil {
		return nil
	}

	if t1 == nil {
		t1 = t2
	}

	if t2 == nil{

	}

	t1.Val = t1.Val + t2.Val
	t1.Left = mergeTrees(t1.Left,t2.Left)
	t1.Right = mergeTrees(t1.Right,t2.Right)

	return t1
}

func preorderTwo(n1,n2 *TreeNode){


	/*

	1. n1 == nil && n2 == nil | return
	2. n1 == nil 以n2为准
	3. n2 == nil 以n1为准
	4. n1 n2 !=nil n1.Val = n1.Val + n2.Val

	preorderTwo(n1.Left,n2.Left)
	preorderTwo(n1.Right,n2.Right)

	*/

}