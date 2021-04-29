package leetcode

func decode(encoded []int, first int) []int {
	var result []int
	result = append(result, first)
	sum := encoded[0] ^ first
	result = append(result, sum)
	for i := 1; i < len(encoded); i++ {
		sum = sum ^ encoded[i]
		result = append(result, sum)
	}
	return result
}
