package binary_tree

import (
	"errors"
)

//根据一棵树的前序遍历与中序遍历构造二叉树。
//
// 注意:
//你可以假设树中没有重复的元素。
//
// 例如，给出
//
// 前序遍历 preorder =[3,9,20,15,7]
//中序遍历 inorder = [9,3,15,20,7]
//
// Related Topics 树 深度优先搜索 数组

func buildTree(preorder []int, inorder []int)*TreeNode{
	return BuildBy(preorder,inorder)
}

func BuildBy(preorder []int, inorder []int)*TreeNode{
	if len(preorder) == 0 || len(inorder) == 0{
		return nil
	}

	node := &TreeNode{
		Val: preorder[0],
	}

	leftPreOrder,leftInOrder,rightPreOrder,rightInOrder := preInOrder(preorder,inorder)

	node.Left = buildTree(leftPreOrder,leftInOrder)
	node.Right = buildTree(rightPreOrder,rightInOrder)

	return node
}


func preInOrder(preorder,inorder []int)([]int,[]int,[]int,[]int){
	//arrLen := len(preorder) // 数组长度

	root := preorder[0]	// 根节点的位置

	// 中序遍历数组，拆分左右
	inOrderRootIdx,_ := arrIdx(inorder,root)
	leftInOrder := inorder[0:inOrderRootIdx]
	rightInOrder := inorder[inOrderRootIdx+1:]

	leftPreOrder := preorder[1:len(leftInOrder)+1]
	rightPreOrder := preorder[len(leftInOrder)+1:]

	return leftPreOrder,leftInOrder,rightPreOrder,rightInOrder
}

func arrIdx(arr []int,needle int)(int,error){
	for k,v := range arr {
		if v == needle{
			return k,nil
		}
	}
	return 0,errors.New("没有找到")
}