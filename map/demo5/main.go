package main

import (
	"fmt"
)

func main() {

	s := "a1c1e1"

	c := []byte(s)
	for i := 1; i < len(s); i = i + 2 {
		c[i] = byte(c[i-1] + c[i] - '0')
	}
	fmt.Println(string(c))
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(root *TreeNode) {
	//前序遍历
	traverse(root.Left)
	//中序遍历
	traverse(root.Right)
	//后序遍历
}
