package link_list

// 拿到节点的前一个节点赋值：node.Next = front
// loop(head.Next) => head==nil head.Next

func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}
