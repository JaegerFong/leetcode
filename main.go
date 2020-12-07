package main

import (
	"fmt"
	"leetcode/binary_tree"
)

//         1
//       / \
//      2   3
//     /   / \
//    4   2   4
//       /
//      4

//    1
//   / \
//  2   3
// / \  /
//4  5 6

// 后续：4,5,2,6,3,1
// 中序：4,2,5,1,6,3

func main() {
	//post := []int{9,15,7,20,3}
	//in := []int{9,3,15,20,7}
	//post := []int{4,2,4,2,4,3,1}
	//in := []int{4,2,1,4,2,3,4}
	post := []int{4, 5, 2, 6, 3, 1}
	in := []int{4, 2, 5, 1, 6, 3}
	tree := binary_tree.BuildByPostAndIn(in, post)
	c := binary_tree.CountNodes(tree)
	fmt.Println(c)
}
