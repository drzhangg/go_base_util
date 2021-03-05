package 数组

func singleNumber(nums []int) int {
	sum := 0
	m := make(map[int]bool)
	for _,v := range nums{
		if _,ok := m[v];ok{
			sum -= v
		}else{
			sum += v
			m[v] = true
		}
	}
	return sum
}
