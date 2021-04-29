package leetcode

func xorOperation(n int, start int) int {
	var result int
	//var arr []int
	for i := 0; i < n; i++ {
		result = result ^(start+2*i)
		//arr = append(arr, ^(start+2*i))
	}
	//fmt.Println(arr)
	return result
}
