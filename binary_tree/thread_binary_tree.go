package binary_tree

/*
不管二叉树的形态如何，空链域的个数总是多过于非空链域的个数，
利用原来的空链域存放指针，指向树中其他结点，这种指针称为线索。

线索化二叉树(以中序遍历为例)：
	1. 如果节点左子树指向为空，则左子树指向（中序遍历中）该结点的前驱结点
	2. 如果节点右子树指向为空，则右子树指向（中序遍历中）该结点的后继结点
另外，还需要增加标志位
	leftTag：用于判断结点指向的是孩子结点还是前驱结点
	rightTag: 用户标志结点指向的是孩子结点还是后继节点
这个过程称为线索化二叉树。
*/

type ThreadBinaryTree struct {
	Val int
	Left *ThreadBinaryTree
	LeftTag int
	Right *ThreadBinaryTree
	RightTag int
}

// 线索化二叉树
var curNode *ThreadBinaryTree
func ThreadingBinaryTree(root *ThreadBinaryTree){
	if root == nil {
		return
	}
	ThreadingBinaryTree(root.Left)

	if root.Left == nil {
		root.LeftTag = 1
		root.Left = root
	}

	if curNode.Right == nil {
		curNode.RightTag = 1
		curNode.Right = root
	}

	curNode = root
	ThreadingBinaryTree(root.Right)
}