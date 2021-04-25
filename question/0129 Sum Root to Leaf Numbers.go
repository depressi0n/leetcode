package question

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var res int
	var dfs func(*TreeNode, []int, int)
	dfs = func(p *TreeNode, s []int, v int) {
		if p.Left == nil && p.Right == nil {
			res += v*10 + p.Val
			return
		}
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
