package leetcode


func kthSmallest(root *TreeNode, k int) int {
	var arr []int
	findArr(root, &arr)
	return arr[k-1]
}

func findArr(root *TreeNode, arr *[]int) {
	if root == nil {
		return
	}

	findArr(root.Left, arr)
	*arr = append(*arr, root.Val)
	findArr(root.Right, arr)
	return
}
