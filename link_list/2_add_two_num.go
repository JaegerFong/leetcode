package link_list

//给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
// 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
// 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
// 示例：
//(2->3) + (1->4->8) = ()
// 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807
//

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curL1, curL2 := l1, l2
	addNum := 0
	list := &ListNode{
		Val:  0,
		Next: nil,
	}
	cur := list
	for curL1 != nil || curL2 != nil || addNum > 0 {

		sumVal := 0
		if curL1 != nil && curL2 != nil {
			sumVal = curL1.Val + curL2.Val
		} else if curL1 != nil {
			sumVal = curL1.Val
		} else if curL2 != nil {
			sumVal = curL2.Val
		}

		if addNum > 0 {
			sumVal += addNum
		}

		if sumVal >= 10 {
			sumVal = sumVal - 10
			addNum = 1
		} else {
			addNum = 0
		}

		cur.Next = &ListNode{
			Val: sumVal,
		}

		cur = cur.Next

		if curL1 != nil {
			curL1 = curL1.Next
		}

		if curL2 != nil {
			curL2 = curL2.Next
		}
	}

	return list.Next
}
