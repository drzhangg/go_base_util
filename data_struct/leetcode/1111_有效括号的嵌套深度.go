package leetcode

func maxDepthAfterSplit(seq string) []int {
	depth := 0
	var ans []int
	for _, c := range seq {
		if c == '(' {
			depth++
			ans = append(ans, depth%2)
		} else {
			ans = append(ans, depth%2)
			depth--
		}
	}
	return ans
}
