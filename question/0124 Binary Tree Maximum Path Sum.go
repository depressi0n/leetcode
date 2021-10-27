package question

import "math"

// 路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。该路径 至少包含一个 节点，且不一定经过根节点。
//路径和 是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
func maxPathSum(root *TreeNode) int {
	return maxPathSumCore(root)
}
// 主要思想：最大路径出现有三种情况，第一种从根节点到叶子节点，第二种是从叶子节点经过根到叶子节点，第三种是叶子节点到叶子节点但不经过根（即出现在子树中
// 使用递归的思想，获得左右子树中对应以上的三种情况对应的值，取其中较大者（利用全局变量保存遍历过程中的最大值）
// 本质上的思想有点像动态规划：从树叶子节点到根节点的动态规划，最后结果是过程中的最大值。
func maxPathSumCore(root *TreeNode) int {
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	maxSum := math.MinInt32
	// 返回值是当前树中从根节点出发的最长路径
	var maxGain func(*TreeNode) int
	maxGain = func(node *TreeNode) int {
		// 如果已经到最底层，则直接返回0
		if node == nil {
			return 0
		}
		// 关键在于定义好调用方式和返回，此外就是怎么更新全局变量
		// 递归计算左右子节点的最大贡献值
		// 只有在最大贡献值大于 0 时，才会选取对应子节点
		leftGain := max(maxGain(node.Left), 0)
		rightGain := max(maxGain(node.Right), 0)

		// 节点的最大路径和取决于该节点的值与该节点的左右子节点的最大贡献值
		priceNewPath := node.Val + leftGain + rightGain

		// 更新答案
		maxSum = max(maxSum, priceNewPath)

		// 返回节点的最大贡献值
		return node.Val + max(leftGain, rightGain)
	}
	maxGain(root)
	return maxSum
}
