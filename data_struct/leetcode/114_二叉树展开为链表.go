package leetcode

func flatten(root *TreeNode) {
	if root == nil {
		return
	}

	flatten(root.Left)
	flatten(root.Right)

	left := root.Left
	right := root.Right

	//将左子树置为空
	root.Left = nil
	// 将左子树移到右子树
	root.Right = left

	//将右子树移到最后
	p := root
	for p.Right != nil {
		p = p.Right
	}
	p.Right = right
}
