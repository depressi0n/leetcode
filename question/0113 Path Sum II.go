package question

func pathSum(root *TreeNode, sum int) [][]int {
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
