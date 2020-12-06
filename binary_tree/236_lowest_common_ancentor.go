package binary_tree

/*
最近公共祖先
递归模型三问：
1.函数用作什么
2.函数参数如何使用
3.函数返回什么数据
*/

/*
1.函数的作用:
判断p，q是否在root，root.Left,root.Right 中


2.函数的参数：
root：当前节点，节点Left,节点Right
p,q 则固定不变


3. 函数返回的数据：
root为null，返回null
p,q都在root中，则返回root；
p,q都不在root中，则返回null；
p在root中，返回q；q在root中，返回q

*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	// p,q在root中
	if left != nil && right != nil {
		return root
	}

	// p,q都不在root中
	if left == nil && right == nil {
		return nil
	}

	if left != nil {
		return left
	}

	return right
}
