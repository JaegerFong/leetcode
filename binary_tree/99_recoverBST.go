package binary_tree

//给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。

var ups []int

func recoverTree(root *TreeNode) {
	ups = []int{}
	inOrderBST(root)
	x, y := swapTow(ups)
	recover(root, 2, x, y)
}

// 中序遍历BST是升序的
func inOrderBST(root *TreeNode) {
	if root == nil {
		return
	}

	inOrderBST(root.Left)
	ups = append(ups, root.Val)
	inOrderBST(root.Right)
}

func swapTow(arr []int) (int, int) {
	x, y := -1, -1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			y = arr[i]
			if x == -1 {
				x = arr[i]
			} else {
				break
			}
		}
	}
	return x, y
}

func recover(root *TreeNode, count, x, y int) {
	if root == nil {
		return
	}

	if root.Val == x || root.Val == y {
		root.Val = y
	} else {
		root.Val = x
	}

	count--
	if count == 0 {
		return
	}

	recover(root.Left, count, x, y)
	recover(root.Right, count, x, y)
}
