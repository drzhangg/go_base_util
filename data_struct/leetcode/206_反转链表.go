package leetcode

func reverseList(head *ListNode) *ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//
	//last := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//return last

	//if head == nil {
	//	return head
	//}

	/*


		var perv *ListNode
		next = head.next
		head.next.next = next


		head.next = head
		head.next.next = head.next

	*/
	var prev *ListNode
	cur, next := head, head
	for cur != nil {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	return prev
}
