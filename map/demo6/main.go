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

}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
