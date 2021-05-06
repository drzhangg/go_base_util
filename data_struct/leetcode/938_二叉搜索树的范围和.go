package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low, high int) int {
	if root == nil {
		return 0
	}

	if root.Val > high {
		return rangeSumBST(root.Left, low, high)
	}

	if root.Val < low {
		return rangeSumBST(root.Right, low, high)
	}

	return root.Val + rangeSumBST(root.Left, low, high) + rangeSumBST(root.Right, low, high)
}

func rangeSumBST1(root *TreeNode, low, high int) int {
	q := []*TreeNode{root}
	var sum int
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			continue
		}

		if node.Val > high {
			q = append(q, node.Left)
		} else if node.Val < low {
			q = append(q, node.Right)
		} else {
			sum += node.Val
			q = append(q, node.Left, node.Right)
		}
	}
	return sum
}
