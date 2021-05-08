package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {

	a := "1010"
	i, _ := strconv.ParseInt(a, 2, 32)
	fmt.Println(i)

	fmt.Println(16 >> 5)
	fmt.Println(16 << 2)

	fmt.Println(34%6)
	fmt.Println(34/6)
	fmt.Println(5&4)
	/*
	0101
	0100
	0100
	0011

	0011
	0010
	0010
	 */


	num := 5
	var count int
	for num>0 {
		num &= num-1
		count++
	}
	fmt.Println(count)

	// 0101  0010
	fmt.Println(5>>1)
	/*
	temp=5>>1
	temp=2
	num=1

	temp=2>>1
	temp=1
	num=num<<1 +1
	num=3

	temp=1>>1
	temp=0
	num=3<<1+1
	num=7

	5 ^ num = 5 ^ 7 = 0101 ^ 0111 = 0010 = 2
	 */
	fmt.Println(5 ^ 7)


}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
