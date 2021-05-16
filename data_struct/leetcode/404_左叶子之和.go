package leetcode

func sumOfLeftLeaves(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var res int
	if root.Left != nil && root.Left.Left == nil && root.Left.Right == nil {
		res += root.Left.Val
	}

	return res + sumOfLeftLeaves(root.Left) + sumOfLeftLeaves(root.Right)
}
