package leetcode

func isPowerOfTwo(n int) bool {
	if n < 0 {
		return false
	}

	if n&(n-1) == 0 {
		return true
	}
	return false
}
/*
4 ->  0100
4-1  ->  0011

0100  &
0011
0000

5 -> 0101
4 -> 0100

0101 &
0100
0100
 */
