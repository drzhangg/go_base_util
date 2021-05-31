package leetcode

var arr []int
func inorderTraversal(root *TreeNode) []int {
	arr = make([]int, 0)
	dfs3(root)
	return arr
}

func dfs3(root *TreeNode) {
	if root != nil {
		dfs3(root.Left)
		arr = append(arr, root.Val)
		dfs3(root.Right)
	}
}

func inorderTraversal1(root *TreeNode) []int {
	//迭代
	//将所有的二叉树节点放到list里
	res := []int{}
	stack := []*TreeNode{}

	for root != nil {
		stack = append(stack, root)
		root = root.Left
	}

	for len(stack) != 0{
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		node = node.Right
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
	}
	return res
}
