package question

//给你一个二叉树的根节点 root ，树中每个节点都存放有一个 0 到 9 之间的数字。
//每条从根节点到叶节点的路径都代表一个数字：
//例如，从根节点到叶节点的路径 1 -> 2 -> 3 表示数字 123 。
//计算从根节点到叶节点生成的 所有数字之和 。
//叶节点 是指没有子节点的节点。
func sumNumbers(root *TreeNode) int {
	return sumNumbersCore(root)
}

// 核心思想：利用DFS遍历从root找到所有的叶子节点，路径上的值就是目标值，维护一个目标值的和
// 递归实现
func sumNumbersCore(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(*TreeNode, []int, int)
	dfs = func(p *TreeNode, s []int, v int) {
		// 叶子节点
		if p.Left == nil && p.Right == nil {
			res += v*10 + p.Val
			return
		}
		// 非叶子节点，更新，s和v可以只维护一个
		s = append(s, p.Val)
		v = v*10 + p.Val
		if p.Left != nil {
			dfs(p.Left, s, v)
		}
		if p.Right != nil {
			dfs(p.Right, s, v)
		}
		s = s[:len(s)-1]
		v = (v - p.Val) / 10
	}
	dfs(root, []int{}, 0)
	return res
}
