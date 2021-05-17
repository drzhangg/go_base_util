package leetcode

func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums,0, len(nums)-1)
}

func build(nums []int, lo, hi int) *TreeNode {
	if lo > hi {
		return nil
	}

	index := lo
	maxVal := nums[index]

	for i := lo +1 ; i <= hi; i++ {
		if maxVal < nums[i] {
			index = i
			maxVal = nums[i]
		}
	}

	root := &TreeNode{Val: maxVal}

	root.Left = build(nums, lo, index-1)
	root.Right = build(nums, index+1, hi)
	return root
}
