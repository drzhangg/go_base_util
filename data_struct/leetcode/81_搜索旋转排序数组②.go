package leetcode

func search2(nums []int, target int) bool {
	l, r := 0, len(nums)-1

	for l != r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return true
		}

		/*
		2,5,6,0,0,2
		l =0
		r = 5
		mid = l + (r - l)/2 = 2

		 */

	}
	return false
}
