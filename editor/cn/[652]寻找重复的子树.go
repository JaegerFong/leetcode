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
// 
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


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type custom struct {
count int
root *TreeNode
}
var childMap map[string]*custom
var repeat []*TreeNode
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
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
//leetcode submit region end(Prohibit modification and deletion)
