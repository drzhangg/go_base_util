package leetcode

func xorQueries(arr []int, queries [][]int) []int {
	var result []int
	for i := 0; i < len(queries); i++ {
		var r int
		for j := queries[i][0]; j < queries[i][1]; j++ {
			r ^= arr[j]
		}
		result = append(result, r)
	}
	return result
}
