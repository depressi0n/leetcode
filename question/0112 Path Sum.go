package question

// 给你二叉树的根节点root 和一个表示目标和的整数targetSum ，判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和targetSum 。
//叶子节点 是指没有子节点的节点。

func hasPathSum(root *TreeNode, sum int) bool {
	return hasPathSumCore(root, sum)
}
func hasPathSumCore(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	//必须是从root到leaf
	var dfs func(*TreeNode, int) bool
	dfs = func(root *TreeNode, cur int) bool {
		if root.Left == nil && root.Right == nil && cur == sum {
			return true
		}
		if (root.Left != nil && dfs(root.Left, cur+root.Left.Val)) || (root.Right != nil && dfs(root.Right, cur+root.Right.Val)) {
			return true
		}
		return false
	}
	return dfs(root, root.Val)
}
