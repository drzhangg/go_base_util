package leetcode

func maxProfit(prices []int) int {
	p, max := prices[0], 0
	for i := 0; i < len(prices); i++ {
		if prices[i] < p {
			p = prices[i]
		}
		if prices[i] - p > max{
			max = prices[i] - p
		}
	}
	return max
}
