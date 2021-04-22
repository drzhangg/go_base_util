package 剑指

type ListNode struct {
	Val int
	Next *ListNode
}

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var newHead *ListNode
	res := []int{}
	for head != nil {
		node := head.Next
		head.Next = newHead
		newHead = head
		head = node
	}

	for newHead != nil {
		res = append(res, newHead.Val)
		newHead = newHead.Next
	}

	return res
}
