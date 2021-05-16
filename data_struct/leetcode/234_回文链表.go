package leetcode

import "fmt"

func isPalindrome(head *ListNode) bool {
	var arr = make([]int,0)
	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	fmt.Println(arr)

	s, e := 0, len(arr)-1

	for s < e {
		fmt.Println("s:",s)
		if arr[s] != arr[e]{
			return false
		}
		s++
		e--
	}

	return true
}
