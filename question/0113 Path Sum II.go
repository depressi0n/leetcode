package question

// 给你二叉树的根节点 root 和一个整数目标和 targetSum ，找出所有 从根节点到叶子节点 路径总和等于给定目标和的路径。
func pathSum(root *TreeNode, sum int) [][]int {
	return pathSumCore(root, sum)
}
func pathSumCore(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	//手动维护一个栈
	res := make([][]int, 0)
	stack := make([]int, 0)
	var dfs func(*TreeNode, int)
	dfs = func(root *TreeNode, curSum int) {
		if root.Left == nil && root.Right == nil && curSum == sum {
			tmp := make([]int, len(stack))
			for i := 0; i < len(stack); i++ {
				tmp[i] = stack[i]
			}
			res = append(res, tmp)
			return
		}
		if root.Left != nil {
			stack = append(stack, root.Left.Val)
			dfs(root.Left, curSum+root.Left.Val)
			stack = stack[:len(stack)-1]
		}
		if root.Right != nil {
			stack = append(stack, root.Right.Val)
			dfs(root.Right, curSum+root.Right.Val)
			stack = stack[:len(stack)-1]
		}
		return
	}
	stack = append(stack, root.Val)
	dfs(root, root.Val)
	return res
}
