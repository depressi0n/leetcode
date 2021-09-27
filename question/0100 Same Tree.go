package question

// 给你两棵二叉树的根节点 p 和 q ，编写一个函数来检验这两棵树是否相同。
//如果两个树在结构上相同，并且节点具有相同的值，则认为它们是相同的。

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return isSameTreeCore(p,q)
}
// 主要思想：递归比较，如果值相同，比较左右孩子，递归出口为nil=nil
func isSameTreeCore(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
