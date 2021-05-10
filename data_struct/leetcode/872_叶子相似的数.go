package leetcode

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var arr1 []int

	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left == nil && node.Right == nil {
			arr1 = append(arr1, node.Val)
			return
		}

		dfs(node.Left)
		dfs(node.Right)
	}
	dfs(root1)

	arrs := append([]int{}, arr1...)
	arr1 = []int{}
	dfs(root2)

	if len(arr1) != len(arrs){
		return false
	}

	for i,v := range arrs{
		if v != arr1[i]{
			return false
		}
	}
	return true
}

func dfs(node *TreeNode) []int {
	var result []int
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		result = append(result, node.Val)
		return nil
	}

	dfs(node.Left)
	dfs(node.Right)
	return result
}
