package main

import (
	"fmt"
	"leetcode/binary_tree"
)

func main(){
	bt := binary_tree.NewBinaryTree([]int{1,0,2,3,5,1})
	fmt.Println(bt.Right.Left.Right)
}