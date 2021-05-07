package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {
	var arr int
	for head != nil {
		arr = (arr << 1) + head.Val
		head = head.Next
	}
	return arr
}
