package leetcode

func bstToGst(root *TreeNode) *TreeNode {
	sum := 0

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}

		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}

	dfs(root)
	return root
}
