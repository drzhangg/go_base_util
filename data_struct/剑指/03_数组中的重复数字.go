package 剑指

/*
输入：
[2, 3, 1, 0, 2, 5, 3]
输出：2 或 3
*/
func findRepeatNumber(nums []int) int {
	if len(nums) < 0 {
		return -1
	}

	var a = make(map[int]int)
	for _, v := range nums {
		if _, ok := a[v]; ok {
			return v
		}
		a[v] = 0
	}
	return -1
}
