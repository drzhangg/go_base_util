package 链表

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	dummy := &ListNode{0,head}

	for node := dummy;node != nil;node = node.Next{
		nodes = append(nodes, node)
	}
	prev := nodes[len(nodes) -1 -n]
	prev.Next = prev.Next.Next
	return dummy.Next
}
