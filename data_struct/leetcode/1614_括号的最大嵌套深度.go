package leetcode

/*
思路：当遇到(时，深度+1，如果当前深度大于max，max等于当前深度，
	遇到)时，深度-1
 */
func maxDepth(s string) int {
	var right, max int
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			right++
			if right > max {
				max = right
			}
		} else if s[i] == ')' {
			right--
		}
	}
	return max
}
