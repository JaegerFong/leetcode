//给定一个二叉树，判断它是否是高度平衡的二叉树。 
//
// 本题中，一棵高度平衡二叉树定义为： 
//
// 
// 一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过 1 。 
// 
//
// 
//
// 示例 1： 
//
// 
//输入：root = [3,9,20,null,null,15,7]
//输出：true
// 
//
// 示例 2： 
//
// 
//输入：root = [1,2,2,3,3,null,null,4,4]
//输出：false
// 
//
// 示例 3： 
//
// 
//输入：root = []
//输出：true
// 
//
// 
//
// 提示： 
//
// 
// 树中的节点数在范围 [0, 5000] 内 
// -104 <= Node.val <= 104 
// 
// Related Topics 树 深度优先搜索 
// 👍 527 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
if postorder(root) >= 0 {
return true
}
return false
}

func postorder(root *TreeNode) int{
if root == nil {
return 0
}

leftDep := postorder(root.Left)
rightDep := postorder(root.Right)

if leftDep >=0 && rightDep >= 0 && math.Abs( float64(leftDep - rightDep)) <=1 {
return max(leftDep,rightDep) + 1
}else{
return -1
}

}

func max(left,right int)int{
if left > right{
return left
}

return right
}
//leetcode submit region end(Prohibit modification and deletion)
