package leetcode

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return dfs4(root.Left,root.Right)
}

func dfs4(root1, root2 *TreeNode) bool {
	if root1 == nil && root2 == nil {
		return true
	}

	if root1 == nil || root2 == nil || root1.Val != root2.Val{
		return false
	}
	return dfs4(root1.Left,root2.Right) && dfs4(root1.Right,root2.Left)
}
