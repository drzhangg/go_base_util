package 链表

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head
	if curr != nil{
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}
