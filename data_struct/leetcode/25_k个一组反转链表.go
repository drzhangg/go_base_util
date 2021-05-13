package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}
	a, b := head, head
	for i := 0; i < k; i++ {
		if b == nil {
			return head
		}
		b = b.Next
	}

	newHead := reverse(a, b)
	a.Next = reverseKGroup(b, k)
	return newHead
}

func reverse(a, b *ListNode) *ListNode {
	var prev *ListNode
	cur, next := a, a
	for cur != b {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}
	return prev
}
