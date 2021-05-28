package leetcode

func isBadVersion(version int) bool {
	return false
}

func firstBadVersion(n int) int {
	if n == 1 && isBadVersion(n) {
		return 1
	}

	left, right := 1, n
	for left < right {
		mid := left + (right-left)/2
		if isBadVersion(mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return left
}
