package main

import (
	"fmt"
	"leetcode/binary_tree"
)

func main(){
	post := []int{9,15,7,20,3}
	in := []int{9,3,15,20,7}
	tree := binary_tree.BuildByPostAndIn(in,post)
	fmt.Println(tree.Right.Left.Val)
}