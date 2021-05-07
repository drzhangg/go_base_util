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

}

func isPowerOfTwo(n int) bool {
	return n > 0 && n&(n-1) == 0
}
