package link_list

// 拿到节点的前一个节点赋值：node.Next = front 1234567890
// loop(head.Next) => head==nil head.Next

// 递归
func reverseList(head *ListNode) *ListNode {
	if head.Next == nil {
		return head
	}

	last := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil

	return last
}

// 迭代  1>2>3>4
func reverseListRecursion(list *ListNode) *ListNode {
	var pre *ListNode = nil
	cur, next := list, list
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}

	return pre
}
