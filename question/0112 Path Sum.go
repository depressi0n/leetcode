package question

func hasPathSum(root *TreeNode, sum int) bool {
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
