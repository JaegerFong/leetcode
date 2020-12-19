package link_list

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {
	list := Generate([]int{1, 2, 3})
	reverseList := reverseList(list.Next)
	fmt.Println(reverseList)
}

func TestReverseRecursion(t *testing.T) {
	list := Generate([]int{1, 2, 3, 4, 5, 6, 7})
	l := list.Next
	reverseList := reverseListRecursion(l)
	fmt.Println(reverseList)
}

func TestReverseNList(t *testing.T) {
	list := Generate([]int{1, 2, 3, 4, 5, 6, 7})
	l := list.Next
	reverseNList := reverseN(l, 3)
	fmt.Println(reverseNList)
}

func TestReverseMNList(t *testing.T) {
	list := Generate([]int{1, 2, 3, 4, 5, 6, 7})
	l := list.Next
	reverseMN := reverseMNList(l, 3, 6)
	fmt.Println(reverseMN)
}

func TestAdd2Num(t *testing.T) {
	l1 := Generate([]int{4, 2, 3})
	l2 := Generate([]int{2, 1})
	l := addTwoNumbers(l1.Next, l2.Next)
	fmt.Println(l)
}
