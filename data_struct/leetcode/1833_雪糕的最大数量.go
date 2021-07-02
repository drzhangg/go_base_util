package leetcode

import "sort"

func maxIceCream(costs []int, coins int) int {
	length := len(costs)
	if length == 0 {
		return 0
	}

	sort.Ints(costs)
	count := 0
	sum := 0
	for i := 0; i < length; i++ {
		if costs[i]+sum <=	 coins {
			count++
			sum += costs[i]
		}
	}
	return count
}
