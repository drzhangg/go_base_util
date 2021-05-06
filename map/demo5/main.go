package main

import (
	"fmt"
)

func main() {

	n := 9
	fmt.Println(n/10)


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
