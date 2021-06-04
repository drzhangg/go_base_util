package leetcode

// 题解：https://leetcode-cn.com/problems/intersection-of-two-linked-lists/solution/xiang-jiao-lian-biao-by-leetcode-solutio-a8jn/
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA== nil || headB == nil {
		return nil
	}
	a,b := headA,headB

	for a!= b{
		if a == nil {
			a = headB
		}else {
			a = a.Next
		}

		if b == nil {
			b = headA
		}else {
			b = b.Next
		}
	}
	return a
}
