package question

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	// res = append(res, []int{root.Val})
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	cur := 0
	end := 1
	for cur < len(queue) {
		if queue[cur].Left != nil {
			queue = append(queue, queue[cur].Left)
		}
		if queue[cur].Right != nil {
			queue = append(queue, queue[cur].Right)
		}

		if cur == end-1 {
			level++
			end = len(queue)
		}
		cur++
	}
	return level
}
