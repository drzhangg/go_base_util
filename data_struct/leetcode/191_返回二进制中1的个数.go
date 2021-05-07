package leetcode

func hammingWeight(n uint32) int {
	var result int
	var mask uint32 = 1
	for i := 0; i < 32; i++ {
		if (n & mask) != 0 {
			result++
		}
		mask = mask << 1
	}
	return result
}
