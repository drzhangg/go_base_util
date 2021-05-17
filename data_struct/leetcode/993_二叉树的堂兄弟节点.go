package leetcode

/*
输入：root = [1,2,3,4], x = 4, y = 3
输出：false

输入：root = [1,2,3,null,4,null,5], x = 5, y = 4
输出：true
 */
func isCousins(root *TreeNode, x int, y int) bool {
	depth = make(map[int]int)
	parent = make(map[int]int)
	dfs2(root,0,0)

	return depth[x] == depth[y] && parent[x] != parent[y]
}

var depth map[int]int
var parent map[int]int

func dfs2(root *TreeNode, par, deep int) {
	if root == nil{
		return
	}

	dfs2(root.Left,root.Val,deep + 1)
	depth[root.Val] = deep
	parent[root.Val] = par
	dfs2(root.Right,root.Val,deep+1)
}
