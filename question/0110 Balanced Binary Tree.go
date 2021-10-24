package question

// 给定一个二叉树，判断它是否是高度平衡的二叉树。

func isBalanced(root *TreeNode) bool {
	return isBalancedCore(root)
}
func isBalancedCore(root *TreeNode) bool {
	// 求树的高度，如果不是平衡树则返回-1
	var height func(root *TreeNode) int
	height = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftHeight := height(root.Left)
		rightHeight := height(root.Right)
		// 如果子树或者本身不是平衡数，直接返回-1
		if leftHeight == -1 || rightHeight == -1 || abs(leftHeight-rightHeight) > 1 {
			return -1
		}
		return max(leftHeight, rightHeight) + 1
	}
	return height(root) >= 0
}
