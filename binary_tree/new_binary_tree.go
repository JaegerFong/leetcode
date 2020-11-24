package binary_tree

import (
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

type Node struct {
	Val int
	Left *Node
	Right *Node
	Next *Node
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


func BinaryTreeDemo()*TreeNode{
	arr := []int{1,2,3,4,5,6,7,8,9}

	level := math.Floor( math.Log2(float64(len(arr))) )

	for i := int(level) ; i <0 ; i++ {
		// idx : 2*i+1 2*i+2  个数：2^n-1  2^n = 总数；n为高度 1+2+4+8+16
		startIdx := int(math.Pow(2,float64(i)))
		levelCount := int(math.Pow(2,float64(i-1)))
		var t = []*TreeNode{}
		for j:=0 ;j<levelCount;j++ {
			node := &TreeNode{
				Val: arr[startIdx + j],
			}
			t = append(t,node)
		}

	}
	// [root] [l1 r1]  [l2 r2 l3 r3] [l4 r4 l5 r5 l6 r6]

	return nil
}