package leetcode

/*
前序遍历 preorder = [3,9,20,15,7]
中序遍历 inorder = [9,3,15,20,7]
*/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	//前序遍历的第一个值，是二叉树的根节点
	root := &TreeNode{preorder[0], nil, nil}
	i := 0
	for ; i < len(preorder); i++ {
		if inorder[i] == preorder[0] {
			break
		}
	}
	//这里i在中序遍历里是根节点的下标，左边是左子树，有边是右子树

	//左子树的长度
	ll := len(inorder[:i])
	//右子树的长度

	//左子树
	root.Left = buildTree(preorder[1:ll+1],inorder[:i])
	root.Right = buildTree(preorder[1+ll:],inorder[i+1:])
	return root
}
