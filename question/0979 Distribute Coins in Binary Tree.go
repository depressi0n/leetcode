package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func DistributeCoins(root *TreeNode) int {
	// 一次可以选择两个相邻节点，从一个节点移动一个coin到另一个节点
	// 要求最后每个节点均匀分布一个coin，而且要求移动次数最少
	// 基本思路，后序遍历每一个节点的val，在每次确定性访问的时候处理栈的信息
	res := 0
	stack := make([]*TreeNode, 0, 100)
	var pre *TreeNode = nil
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		// 上面一定会走到p==nil的情况，就是左子树必定没有
		p = stack[len(stack)-1]
		// 本应该在这里访问p
		stack = stack[:len(stack)-1]
		if p.Right == nil || p.Right == pre {
			if p.Val > 1 {
				res += p.Val - 1
				stack[len(stack)-1].Val += p.Val - 1

			} else if p.Val < 1 {
				res += 1 - p.Val
				stack[len(stack)-1].Val -= 1 - p.Val
			}
			p.Val = 1
			pre = p
			p = nil
		} else {
			stack = append(stack, p)
			p = p.Right
		}

	}
	return res
	// 树形dp？感觉可以考虑一下？TODO
	// TODO 递归的思想是怎么想进去的...
}
