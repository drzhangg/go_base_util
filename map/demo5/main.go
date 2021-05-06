package main

import "fmt"

func main() {
	word1 := []string{"abc", "d", "defg"}
	word2 := []string{"abcddefg"}
	var b1, b2 []byte
	for i := 0; i < len(word1); i++ {
		for _, v := range word1[i] {
			b1 = append(b1, byte(v))
		}
	}
	for i := 0; i < len(word2); i++ {
		for _, v := range word2[i] {
			b2 = append(b2, byte(v))
		}
	}
	fmt.Println(string(b1))
	fmt.Println(string(b2))
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
