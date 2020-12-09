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
