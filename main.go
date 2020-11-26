package main

import (
	"fmt"
	"leetcode/binary_tree"
)

func main(){
	//post := []int{9,15,7,20,3}
	//in := []int{9,3,15,20,7}
	post := []int{4,2,4,2,4,3,1}
	in := []int{4,2,1,4,2,3,4}
	tree := binary_tree.BuildByPostAndIn(in,post)

	fmt.Println(binary_tree.DfsDep(tree))
}