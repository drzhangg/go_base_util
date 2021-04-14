package 二维数组

func rotate(matrix [][]int)  {
	for i := 0; i < len(matrix)/2; i++ {
		temp := matrix[i]
		matrix[i] = matrix[len(matrix) -i - 1]
		matrix[len(matrix) - i - 1] = temp
	}

	for i := 0; i < len(matrix); i++ {
		for j := i+1; j < len(matrix); j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}
