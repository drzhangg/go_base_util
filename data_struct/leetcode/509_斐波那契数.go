package leetcode

func fib1(n int) int {
	if n <= 1 {
		return n
	}

	var dp [2]int
	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {
		sum := dp[0] + dp[1]
		dp[0] = dp[1]
		dp[1] = sum
	}
	return dp[1]
}
