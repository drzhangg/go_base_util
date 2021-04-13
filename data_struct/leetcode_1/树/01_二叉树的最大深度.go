package 树

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var queue []*TreeNode       //辅助队列
	queue = append(queue, root) //根节点入队
	depth := 0                  //初始化树深度为0

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			s := queue[0]      //获取队首元素
			queue = queue[1:]  //队首元素出队
			if s.Left != nil { //如果左子树不为空，左子树入队
				queue = append(queue, s.Left)
			}
			if s.Right != nil { //如果右子树不为空，右子树入队
				queue = append(queue, s.Right)
			}
		}
		depth++
	}
	return depth
}
