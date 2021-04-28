package leetcode

func maximumWealth(accounts [][]int) int {
	if accounts == nil {
		return 0
	}

	//var result []int
	var max = 0
	for i := 0; i < len(accounts); i++ {
		var sum int
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if max < sum{
			max = sum
		}
		//result = append(result, sum)
	}
	//sort.Ints(result)


	return max
}
