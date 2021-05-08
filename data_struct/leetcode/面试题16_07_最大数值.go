package leetcode

func maximum(a int, b int) int {
	ret := int64(a-b)
	ret = int64(a) - ret&(ret>>63)
	return int(ret)
}
/*

a=3,b=1
a-b=2
>>63 0000
0010 & 0000 = 0000
a - 0 = a

a=1,b=3
a-b= -2
>>63 1111
1110 & 1111 = 1110
a - 1110 = 1 - (-2) = 3


 */