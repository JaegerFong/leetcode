//给定一个不含重复元素的整数数组。一个以此数组构建的最大二叉树定义如下： 
//
// 
// 二叉树的根是数组中的最大元素。 
// 左子树是通过数组中最大值左边部分构造出的最大二叉树。 
// 右子树是通过数组中最大值右边部分构造出的最大二叉树。 
// 
//
// 通过给定的数组构建最大二叉树，并且输出这个树的根节点。 
//
// 
//
// 示例 ： 
//
// 输入：[3,2,1,6,0,5]
//输出：返回下面这棵树的根节点：
//
//      6
//    /   \
//   3     5
//    \    / 
//     2  0   
//       \
//        1
// 
//
// 
//
// 提示： 
//
// 
// 给定的数组的大小在 [1, 1000] 之间。 
// 
// Related Topics 树 
// 👍 216 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
return build(nums,0,len(nums)-1)
}


func build(nums []int,start,end int) *TreeNode{
if start>end {
return nil
}
// 最大的值
mid,val := maxValAndIdx(nums[start:end])

return &TreeNode{
Val: val,
Left: build(nums,start,mid),
Right: build(nums,mid+1,end),
}
}

func maxValAndIdx(nums []int)(max,idx int){
idx = 0
max = nums[0]
for k, v := range nums {
if v > max {
max = v
idx = k
}
}

return max,idx
}

//leetcode submit region end(Prohibit modification and deletion)
