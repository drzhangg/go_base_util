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
		1.next = 2
		2.next = 3

		->
		1.next.next =3
		3.next = 2
		3.next = 2
		2.next = 1
	*/
	var prev *ListNode
	temp := head
	for temp != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return prev
}
