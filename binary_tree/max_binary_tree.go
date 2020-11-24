package binary_tree

// 最大二叉树
// 二叉树的根是数组中的最大元素。
// 左子树是通过数组中最大值左边部分构造出的最大二叉树。
// 右子树是通过数组中最大值右边部分构造出的最大二叉树。
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
// [3,2,1] [6] [0,5]
// [3] [2] [1] [6] [0] [5]
func ConstructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums,0,len(nums)-1)
}


func build(nums []int,start,end int) *TreeNode{
	if start > end {
		return nil
	}
	// 最大的值
	maxVal,idx := nums[start],start
	for i:=start ; i<end ; i ++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			idx = i
		}
	}

	return &TreeNode{
		Val: maxVal,
		Left: build(nums,start,idx),
		Right: build(nums,idx+1,end),
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
