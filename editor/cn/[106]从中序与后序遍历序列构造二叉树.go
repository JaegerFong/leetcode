//根据一棵树的中序遍历与后序遍历构造二叉树。 
//
// 注意: 
//你可以假设树中没有重复的元素。 
//
// 例如，给出 
//
// 中序遍历 inorder = [9,3,15,20,7]
//后序遍历 postorder = [9,15,7,20,3] 
//
// 返回如下的二叉树： 
//
//     3
//   / \
//  9  20
//    /  \
//   15   7
// 
// Related Topics 树 深度优先搜索 数组 
// 👍 406 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(inorder []int, postorder []int) *TreeNode {
	return BuildByPostAndIn(inorder,postorder)
}

func BuildByPostAndIn(inorder []int, postorder []int) *TreeNode {
if len(inorder) == 0 || len(postorder) == 0 {
return nil
}

node := &TreeNode{
Val: postorder[len(postorder) -1 ],
}

leftInorder,leftPostorder,rightInorder,rightPostorder := postInOrder(inorder,postorder)

node.Left = BuildByPostAndIn(leftInorder,leftPostorder)
node.Right = BuildByPostAndIn(rightInorder,rightPostorder)

return node
}

// 通过后续和中序得出各自左右节点的后续，中序数据
func postInOrder(inorder ,postorder []int)([]int,[]int,[]int,[]int){
arrLen := len(postorder)
root := postorder[arrLen-1]

inorderRootIdx,_ := arrIdx(inorder,root)
leftInOrder := inorder[0:inorderRootIdx]
rightInOrder := inorder[inorderRootIdx+1:]

leftInOrderLen := len(leftInOrder)
leftPostorder := postorder[0:leftInOrderLen]
rightPostorder := postorder[leftInOrderLen:arrLen-1]

return leftInOrder,leftPostorder,rightInOrder,rightPostorder

}

func arrIdx(arr []int,needle int)(int,error){
for k,v := range arr {
if v == needle{
return k,nil
}
}
return 0,errors.New("没有找到")
}
//leetcode submit region end(Prohibit modification and deletion)
