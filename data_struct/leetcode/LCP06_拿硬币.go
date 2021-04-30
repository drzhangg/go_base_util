package leetcode

func minCount(coins []int) int {
	var res int
	for i := 0; i < len(coins); i++ {
		var r int
		if coins[i] %2 == 0{
			r = coins[i] /2
		}else {
			r = coins[i] /2 +1
		}
		res += r
	}
	return res
}
