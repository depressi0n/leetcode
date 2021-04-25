package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	flag := true
	//用递归会很方便
	if root.Left != nil {
		//左子树最大节点
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		flag = p.Val < root.Val && isValidBST(root.Left)
	}
	if flag {
		if root.Right != nil {
			//右子树最小节点
			p := root.Right
			for p.Left != nil {
				p = p.Left
			}
			if p.Val > root.Val && isValidBST(root.Right) {
				return true
			}
			return false
		}
		return true
	}
	return false
}
