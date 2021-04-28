package leetcode

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var max = 0
	for i := 0; i < len(candies); i++ {
		if max < candies[i]{
			max = candies[i]
		}
	}

	var result [] bool
	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= max {
			result = append(result, true)
		}else {
			result = append(result, false)
		}
	}
	return result
}
