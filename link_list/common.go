package link_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func Generate(arr []int) *ListNode {
	list := &ListNode{
		Val: 0,
	}
	cur := list
	for _, v := range arr {
		cur.Next = &ListNode{
			Val: v,
		}
		cur = cur.Next
	}

	return list
}
