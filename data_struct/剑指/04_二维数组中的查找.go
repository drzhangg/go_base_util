package 剑指

func findNumberIn2DArray(matrix [][]int, target int) bool {
	var i = len(matrix) - 1
	var j = 0
	for i >= 0 && j < len(matrix[0]) {
		if matrix[i][j] > target {
			i--
		} else if matrix[i][j] < target {
			j++
		} else {
			return true
		}
	}
	return false
}
