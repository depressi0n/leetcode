package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func navigation(root *TreeNode) int {
	// 从N个节点中选择M节点
	// 计算其他节点到这M个节点的距离，保证任两个节点的距离向量均不相同
	// 思路：将二叉树按照图处理，计算距离，得到一个距离矩阵，选择其中M列，保证剩下的行没有相同的两行
	// 上述思路是做不下去的

	// 证明：如果树不是一条链，在最优化导航装置的位置的前提下，能够区分树中所有节点的 充分必要条件是：
	// 所有三叉节点的 3 个子树中，至少有2 个子树中有导航装置。
	//TODO:看懂
	res, s := 0, 1
	var dfs func(node *TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		l := dfs(node.Left)
		r := dfs(node.Right)
		if node.Left != nil && node.Right != nil {
			if !l && !r {
				res += 1
			}
			if !(l && r) {
				s = 1
			} else {
				s = 0
			}
			return true
		}
		return l || r
	}
	l := dfs(root.Left)
	r := dfs(root.Right)
	if l && r {
		return res
	}
	return res + s
}
