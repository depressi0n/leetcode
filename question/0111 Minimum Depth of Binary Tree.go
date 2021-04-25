package question

import "math"

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	min := math.MaxInt32
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, cur int) {
		if root.Left == nil && root.Right == nil {
			if cur < min {
				min = cur
			}
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
	return min
}

func minDepth_1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}
	min := math.MaxInt32
	if root.Left != nil {
		sub := minDepth(root.Left)
		if sub < min {
			min = sub + 1
		}
	}
	if root.Right != nil {
		sub := minDepth(root.Right)
		if sub < min {
			min = sub + 1
		}
	}
	return min
}
