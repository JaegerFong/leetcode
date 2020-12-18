package link_list

// 反转链表的前n个节点
var head *ListNode

func reverseN(list *ListNode, n int) *ListNode {
	if n == 1 {
		head = list.Next
		return list
	}

	last := reverseN(list.Next, n-1)
	list.Next.Next = list
	list.Next = head
	return last
}
