package question

// 给你一个二叉树的根节点 root ，判断其是否是一个有效的二叉搜索树。
//有效 二叉搜索树定义如下：
//		节点的左子树只包含 小于 当前节点的数。
//		节点的右子树只包含 大于 当前节点的数。
//		所有左子树和右子树自身必须也是二叉搜索树。
func isValidBST(root *TreeNode) bool {
	return isValidBSTCore(root)
}

// 首先是二叉搜索书，且对子树也要求是二叉搜索树
// 最简单对方法是用递归：递归出口定义为nil，如果左子树不空，则判断左子树是否满足二叉搜索树的定义以及左孩子的值是否小于根的值，对于右孩子同理
func isValidBSTCore(root *TreeNode) bool {
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
		flag = p.Val < root.Val && isValidBSTCore(root.Left)
	}
	if flag {
		if root.Right != nil {
			//右子树最小节点
			p := root.Right
			for p.Left != nil {
				p = p.Left
			}
			if p.Val > root.Val && isValidBSTCore(root.Right) {
				return true
			}
			return false
		}
		return true
	}
	return false
}
