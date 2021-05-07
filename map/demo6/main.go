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
	i,_ := strconv.ParseInt(a,2,32)
	fmt.Println(i)

}
