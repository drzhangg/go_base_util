package leetcode

func hammingDistance(x int, y int) int {
	num := x ^ y
	var count int
	for num != 0 {
		num = num & (num-1)
		count++
	}
	return count

	//var count int
	//for x != y {
	//	if x > y {
	//		y <<= 1
	//		count++
	//	} else {
	//		x <<= 1
	//		count++
	//	}
	//}
	//return count
}
