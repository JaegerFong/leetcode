package binary_tree

import (
	"strconv"
)

//给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//
// 两棵树重复是指它们具有相同的结构以及相同的结点值。
//
// 示例 1：
//
//         1
//       / \
//      2   3
//     /   / \
//    4   2   4
//       /
//      4
// 中序 4,2,1,4,2,3,4
// 后序 4,2,4,2,4,3,1
//
// 下面是两个重复的子树：
//
//       2
//     /
//    4
//
//
// 和
//
//     4
//
//
// 因此，你需要以列表的形式返回上述重复子树的根结点。
// Related Topics 树
// 👍 180 👎 0

/*
思路：
1、确定以某个节点为根的整颗树结构
2、判断当前节点和其他节点是否一致
*/

type custom struct {
	count int
	root *TreeNode
}
var childMap map[string]*custom
var repeat []*TreeNode
func FindDuplicateSubtrees(root *TreeNode) []*TreeNode {
	childMap = make(map[string]*custom ,100)
	repeat = nil
	if root == nil {
		return nil
	}

	PostOrder(root)
	for _,v := range childMap {
		if v.count > 1 {
			repeat = append(repeat,v.root)
		}
	}
	return repeat
}

func PostOrder(root *TreeNode)string{
	if root == nil {
		return "#"
	}

	left := PostOrder(root.Left)
	right := PostOrder(root.Right)
	val := root.Val

	s := left + "," + right + "," + strconv.Itoa(val)
	if _,ok := childMap[s];!ok {
		// 不存在，则保存节点
		childMap[s] = &custom{
			count: 1,
			root:root,
		}
	}else{
		// 已经存在，则说明重复了
		childMap[s].count += 1
	}

	return s
}