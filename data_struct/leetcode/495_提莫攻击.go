package leetcode

// [1,4] 2   return 4
func findPoisonedDuration(timeSeries []int, duration int) int {

	if len(timeSeries) == 0 {
		return 0
	}

	var result int
	for i := 1; i < len(timeSeries); i++ {
		gap := timeSeries[i] - timeSeries[i-1]
		if gap > duration {
			result += duration
		} else {
			result += gap
		}
	}
	return result + duration
}
