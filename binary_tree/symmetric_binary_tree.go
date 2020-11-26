package binary_tree
//请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

// 左子树的左节点等于右子树的右节点
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return Symmetric(root.Left,root.Right)
}


func Symmetric(leftNode , rightNode *TreeNode) bool{
	if leftNode == nil && rightNode == nil {
		return true
	}

	if (leftNode == nil && rightNode !=nil ) ||
		(leftNode != nil && rightNode == nil ) ||
		(leftNode.Val != rightNode.Val){
		return false
	}

	return Symmetric(leftNode.Left,rightNode.Right) && Symmetric(leftNode.Right,rightNode.Left)
}