package question

// 给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数。
//说明: 叶子节点是指没有子节点的节点。

func maxDepth(root *TreeNode) int {
	return maxDepthCore3(root)
}

// 层次遍历
func maxDepthCore1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 0
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

// 利用从根节点到叶子节点到路径长度来确定最大深度，只能使用后序遍历
func maxDepthCore2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	depth := 1
	q := make([]*TreeNode, 0, 100)
	p := root
	var pre *TreeNode
	for len(q) != 0 || p != nil {
		for p != nil {
			q = append(q, p)
			p = p.Left
		}
		p = q[len(q)-1]
		if p.Right != nil && p.Right != pre {
			p = p.Right
		} else {
			q = q[:len(q)-1]
			pre = p
			p = nil
			depth = max(depth, len(q)+1)
		}
	}
	return depth
}

// 递归，使用深度优先搜索 depth(root) = max(depth(root.Left),depth(root.Right))+1
func maxDepthCore3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepthCore3(root.Left), maxDepthCore3(root.Right)) + 1
}
