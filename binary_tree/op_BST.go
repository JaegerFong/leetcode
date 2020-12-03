package binary_tree

func insertIntoBST(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return &TreeNode{
			Val: val,
		}
	}

	if val < node.Val {
		node.Left = insertIntoBST(node.Left, val)
	}

	if val > node.Val {
		node.Right = insertIntoBST(node.Right, val)
	}

	return node
}

/*
node.Val == val ,找到了进行处理
1. 没有左右子树，则直接删除
2. 左子树为空，则返回右子树；右子树为空，则返回左子树
3. 左右子树都不为空，为了不破坏BST特性，需要知道接替者，找左子树最大的节点或者右子树最小的节点来代替
val > node.Val ,去右子树找
val < node.Val ，去左子树找
*/
func deleteNode(node *TreeNode, val int) *TreeNode {
	if node == nil {
		return nil
	}

	if node.Val == val {
		if node.Left == nil && node.Right == nil {
			return nil
		} else if node.Left == nil {
			return node.Right
		} else if node.Right == nil {
			return node.Left
		} else {
			minNode := minNodeInBST(node.Right)
			node.Val = minNode.Val
			node.Right = deleteNode(node.Right, minNode.Val)
		}
	} else if val > node.Val {
		deleteNode(node.Left, val)
	} else if val < node.Val {
		deleteNode(node.Right, val)
	}

	return node
}

func minNodeInBST(node *TreeNode) *TreeNode {
	for node.Left != nil {
		node = node.Left
	}

	return node
}
