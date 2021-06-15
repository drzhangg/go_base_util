package leetcode

func longestPalindrome(s string) string {
	length := len(s)

	if length <=1 {
		return s
	}

	dp := make([][]bool,length)

	start := 0
	maxlen := 1

	for i := 0; i < length; i++ {
		dp[i] = make([]bool,length)
		dp[i][i] = true
		for j := 0; j < i; j++ {
			if s[i] == s[j] && (i - j <=2||dp[j+1][i-1]){
				dp[j][i]=true
			}else {
				dp[j][i]=false
			}

			if dp[j][i] {
				curlen := i-j+1
				if curlen > maxlen {
					maxlen = curlen
					start = j
				}
			}
		}
	}
	return s[start:start+maxlen]
}
