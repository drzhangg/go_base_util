package leetcode

func fib1(n int) int {
	if n <= 1 {
		return n
	}

	p, q, r := 0, 0, 1
	for i := 2; i < n; i++ {
		p = q
		q=r
		r = p + q
	}
	return r	
}

