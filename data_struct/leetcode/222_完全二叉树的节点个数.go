package leetcode

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	res := left + right + 1
	return res
}
