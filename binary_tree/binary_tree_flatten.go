package binary_tree

// 例如，给定二叉树
//
//     1
//   / \
//  2   5
// / \   \
//3   4   6
//
// 将其展开为：
//
// 1
// \
//  2
//   \
//    3
//     \
//      4
//       \
//        5
//         \
//          6

// 二叉树展开为链表
/*
1.
*/
func flatten(root *TreeNode)  {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	root.Left = nil
	root.Right = left

	// p = p right
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right =right
}