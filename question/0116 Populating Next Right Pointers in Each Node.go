package question

// 给定一个 完美二叉树 ，其所有叶子节点都在同一层，每个父节点都有两个子节点。
// 填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
// 初始状态下，所有next 指针都被设置为 NULL。

// 进阶：
//你只能使用常量级额外空间。
//使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

// 层次遍历也可，但需要存储空间，所以这里直接用递归
func connect116(root *Node116) *Node116 {
	return connect116Core(root)
}
func connect116Core(root *Node116) *Node116 {
	// 这是一棵完全二叉树
	// 递归
	// 如果节点为空 或 节点是叶子节点 则直接返回
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	// it means that it is wrongly
	// 说明有错误，因为必然是两个节点均有的情况
	if root.Left == nil || root.Right == nil {
		return nil
	}
	// 递归处理左子树
	connect116Core(root.Left)
	// 递归处理右子树
	connect116Core(root.Right)
	// 将左子树和右子树连接起来，使用左子树的右孩子连接右子树的左孩子
	p := root.Left
	q := root.Right
	p.Next = q
	// 只要有一个为空，就不用往下走，因为下面必然已经是末尾了
	for p.Right != nil && q.Left != nil {
		p = p.Right
		q = q.Left
		p.Next = q
	}
	return root
}
