package leetcode

/*
只有一个链表为空时，需要继续处理，初始化nil为0

*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{0, nil}
	tmp := result
	more := 0

	for l1 != nil || l2 != nil {
		if l1 == nil {
			l1 = &ListNode{0, nil}
		}

		if l2 == nil {
			l2 = &ListNode{0, nil}
		}

		sum := l1.Val + l2.Val + more
		tmp.Val = sum % 10

		if sum >= 10 {
			more = 1
		} else {
			more = 0
		}

		l1 = l1.Next
		l2 = l2.Next

		if l1 == nil && l2 == nil {
			if more == 1{
				tmp.Next = new(ListNode)
				tmp = tmp.Next
				tmp.Val = 1
			}
			break
		}

		tmp.Next = new(ListNode)
		tmp = tmp.Next

	}
	return result
}
