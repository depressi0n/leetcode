package question

import "math"

// 给定一个二叉树，找出其最小深度。
//最小深度是从根节点到最近叶子节点的最短路径上的节点数量。
//说明：叶子节点是指没有子节点的节点。

func minDepth(root *TreeNode) int {
	return minDepthCore(root)
}

// DFS
func minDepthCore(root *TreeNode) int {
	if root == nil {
		return 0
	}
	minHeight := math.MaxInt32
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cur int) {
		if root.Left == nil && root.Right == nil {
			if cur < minHeight {
				minHeight = cur
			}
			return
		}
		// 剪枝，如果当前深度已经比见到的最小深度要大，答案不可能在其中
		if cur >= minHeight {
			return
		}
		if root.Left != nil {
			dfs(root.Left, cur+1)
		}
		if root.Right != nil {
			dfs(root.Right, cur+1)
		}
	}
	dfs(root, 1)
	return minHeight
}

// 使用最小高度的递归方式
func minDepthCore2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	minHeight := math.MaxInt32
	if root.Left != nil {
		sub := minDepthCore2(root.Left)
		if sub < minHeight {
			minHeight = sub + 1
		}
	}
	if root.Right != nil {
		sub := minDepthCore2(root.Right)
		if sub < minHeight {
			minHeight = sub + 1
		}
	}
	return minHeight
}

// 层次遍历即BFS
func minDepthCore3(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := make([]*TreeNode, 0, 100)
	q = append(q, root)
	cur, end := 0, 1
	depth := 1
	for cur < len(q) {
		if q[cur].Left == nil && q[cur].Right == nil {
			break
		}
		if q[cur].Left != nil {
			q = append(q, q[cur].Left)
		}
		if q[cur].Right != nil {
			q = append(q, q[cur].Right)
		}
		cur++
		if cur == end {
			depth++
			end = len(q)
		}
	}
	return depth
}
