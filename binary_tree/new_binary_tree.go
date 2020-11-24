package binary_tree

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

// 求个二叉树层级：level := math.Floor( math.Log2(float64(len(arr))) )

// TODO 修改
func NewBinaryTree(arr []int)(root *TreeNode){
	node := createOrder(0,arr)
	return node
}

func createOrder(idx int,arr []int)*TreeNode{
	node := &TreeNode{}
	if idx > len(arr)-1{
		return node
	}

	if arr[idx] == 0 {
		return node
	}

	node.Val = arr[idx]
	node.Left = createOrder(idx*2+1,arr)
	node.Right = createOrder(idx*2+2,arr)
	return node
}