package leetcode

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val < key{
		root.Right = deleteNode(root.Right,key)
	}else if root.Val > key{
		root.Left = deleteNode(root.Left,key)
	}else { //root的val 等于 要删除的节点
		if root.Right == nil {
			return root.Left
		}

		if root.Left == nil{
			return root.Right
		}

		//找到右子树的最小节点
		min := root.Right
		for min.Left != nil {
			min = min.Left
		}

		min.Left = root.Left
		root = root.Right
	}
	return root
}