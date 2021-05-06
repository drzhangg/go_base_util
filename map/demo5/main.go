package main

func main() {

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
