package main

import (
	"go_base_util/data_struct/leetcode_1/链表"
)

func main() {

	/*
		3
		2
		1
	*/
}

/*
输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
*/
func reverseList(head *链表.ListNode) *链表.ListNode {
	//if head == nil || head.Next == nil {
	//	return head
	//}
	//
	//p := reverseList(head.Next)
	//head.Next.Next = head
	//head.Next = nil
	//return p

	var pre *链表.ListNode

	cur := head
	for head != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
