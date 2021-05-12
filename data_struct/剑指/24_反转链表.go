package 剑指

func reverseList(head *ListNode) *ListNode {

	var prev *ListNode
	temp := head
	for head != nil {
		next := temp.Next
		temp.Next = prev
		prev = temp
		temp = next
	}
	return prev
}
