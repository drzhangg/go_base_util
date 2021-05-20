package 剑指

func fib(n int)int{

	constant := 1000000007
	first := 0
	second := 1
	for n>0{
		temp := first + second
		first = second % constant
		second = temp % constant
		n--
	}
	return first
}

func fib2(n int) int {
	if n ==0{
		return 0
	}
	dp := make([]int,n+1)
	dp[0] ,dp[1]= 0,1
	for i:=2;i <=n;i++{
		dp[i] = dp[i-1]%1000000007 +dp[i-2]%1000000007
	}
	return dp[n]%1000000007
}

