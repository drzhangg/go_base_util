package leetcode

import "sort"

func searchMatrix(matrix [][]int, target int) bool {
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}
	return false
}

func searchMatrix2(matrix [][]int, target int) bool {
	row := sort.Search(len(matrix), func(i int) bool {
		return matrix[i][0] > target
	}) - 1

	if row < 0 {
		return false
	}

	col := sort.SearchInts(matrix[row], target)
	return col < len(matrix[row]) && matrix[row][col] == target
}

func searchMatrix3(matrix [][]int, target int) bool {
	m,n := len(matrix),len(matrix[0])
	i := sort.Search(m*n , func(i int) bool {
		return matrix[i/n][i%n] >=target
	})
	return i< m*n &&matrix[i/n][i%n]==target
}
