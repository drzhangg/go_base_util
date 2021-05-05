package leetcode

import "math"

func minTimeToVisitAllPoints(points [][]int) int {
	result := 0
	for i := 1; i < len(points); i++ {
		result += int(math.Max(math.Abs(float64(points[i][0]-points[i-1][0])),math.Abs(float64(points[i][1] - points[i-1][1]))))
	}
	return result
}
