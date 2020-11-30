package binary_tree

//给定一个二叉搜索树，编写一个函数 kthSmallest 来查找其中第 k 个最小的元素。
// 二叉搜索树的一个重要特性：中序遍历得到的数列是有序的（升序）
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return 0
	}

	arr = nil
	inorder_bst(root)

	// 找k位置
	target := 0
	for idx, val := range arr {
		if idx+1 == k {
			target = val
		}
	}

	return target
}

var arr []int

func inorder_bst(root *TreeNode) {
	if root == nil {
		return
	}

	inorder_bst(root.Left)
	arr = append(arr, root.Val)
	inorder_bst(root.Right)
}
