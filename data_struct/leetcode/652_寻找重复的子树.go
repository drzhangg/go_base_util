package leetcode

import "strconv"

func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	myMap := make(map[string]int)
	ret := make([]*TreeNode, 0)
	dfs1(root, myMap, &ret)
	return ret
}

func dfs1(root *TreeNode, mayMap map[string]int, ret *[]*TreeNode) string {
	if root == nil {
		return ""
	}

	left := dfs1(root.Left, mayMap, ret)
	right := dfs1(root.Right, mayMap, ret)

	retString := left + "," + right + "," + strconv.Itoa(root.Val)

	if v, ok := mayMap[retString]; !ok {
		mayMap[retString] = 1
	} else {
		if v == 1 {
			*ret = append(*ret, root)
		}
		mayMap[retString]++
	}
	return retString
}
