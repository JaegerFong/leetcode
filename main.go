package main

import (
	"fmt"
	"leetcode/binary_tree"
)

func main(){
	pre := []int{3,9,20,15,7}
	in := []int{9,3,15,20,7}
	tree := binary_tree.BuildBy(pre,in)
	fmt.Println(tree.Right.Left.Val)
}