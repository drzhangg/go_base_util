package leetcode

func checkIfPangram(sentence string) bool {
	//if len(sentence) < 26{
	//	return false
	//}
	//var m = make(map[int32]int)
	//for _,v := range sentence{
	//	m[v]++
	//}
	//return len(m) == 26

	//大佬解法
	m := 0
	for _,v := range sentence{
		m |= 1 << (v - 'a')
	}
	return m == 1 << 26 -1
}
