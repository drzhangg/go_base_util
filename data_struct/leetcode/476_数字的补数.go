package leetcode

func findComplement(num int) int {
	var count int
	temp := num
	for temp > 0 {
		temp>>= 1
		count = (count<<1)+1
	}
	return count ^ num
}
