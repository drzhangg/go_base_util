package leetcode

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect1(root *Node) *Node {
	if root == nil {
		return nil
	}
	connectTwoNode(root.Left,root.Right)
	return root
}

func connectTwoNode(node1 *Node, node2 *Node) {
	if node1 == nil || node2 == nil {
		return
	}

	node1.Next = node2

	//连接同一父节点的左右节点
	connectTwoNode(node1.Left, node1.Right)
	connectTwoNode(node2.Left, node2.Right)

	// 连接不同父节点的右节点和左节点
	connectTwoNode(node1.Right, node2.Left)
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	queue := []*Node{root}

	// 循环迭代二叉树层数
	for len(queue) > 0 {
		tmp := queue
		queue = nil

		for i, node := range tmp {
			// 连接
			if i+1 < len(tmp) {
				node.Next = tmp[i+1]
			}

			// 拓展下一层
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return root
}
