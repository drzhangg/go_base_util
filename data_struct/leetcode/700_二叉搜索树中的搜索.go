package leetcode

func searchBST(root *TreeNode, val int) *TreeNode {
	//递归
	//if root == nil {
	//	return nil
	//}
	//
	//if root.Val > val{
	//	return searchBST(root.Left,val)
	//}else if root.Val < val{
	//	return searchBST(root.Right,val)
	//}else {
	//	return root
	//}

	//迭代
	for root !=nil{
		if root.Val > val {
			root = root.Left
		}else if root.Val < val{
			root = root.Right
		}else {
			return root
		}
	}
	return nil
}
