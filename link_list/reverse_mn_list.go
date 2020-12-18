package link_list

func reverseMNList(list *ListNode, m, n int) *ListNode {
	if m == 1 {
		return reverseN(list, n)
	}

	list.Next = reverseMNList(list.Next, m-1, n-1)
	return list
}
