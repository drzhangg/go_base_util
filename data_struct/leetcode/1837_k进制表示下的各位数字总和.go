package leetcode

func sumBase(n int, k int) int {
	var result int
	for n > 0 {
		result+= n%k
		n/=k
	}
	return result
}
