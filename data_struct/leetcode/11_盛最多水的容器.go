package leetcode

/*
双指针
头指针指向数组头结点
尾指针指向数组尾节点
数组和 = 头尾节点的较小值 * 数组的长度
返回最大的数组和
*/
func maxArea(height []int) int {
	s, e := 0, len(height)-1
	result := 0

	for s < e {
		area := getMain(height[s], height[e])*(e-s)
		result = getMax(area,result)
		if height[s] <= height[e]{
			s++
		}else {
			e--
		}
	}
	return result
}

func getMain(a, b int) int {
	if a > b {
		return b
	}
	return a
}



