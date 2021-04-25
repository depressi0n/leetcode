package question

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Left *Node
 *     Right *Node
 *     Next *Node
 * }
 */

type Node116 struct {
	Val   int
	Left  *Node116
	Right *Node116
	Next  *Node116
}

func connect(root *Node116) *Node116 {
	// 这是一棵完全二叉树
	// 递归
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	// it means that it is wrongly
	if root.Left == nil || root.Right == nil {
		return nil
	}
	connect(root.Left)
	connect(root.Right)
	p := root.Left
	q := root.Right
	p.Next = q
	for p.Right != nil && q.Left != nil {
		p = p.Right
		q = q.Left
		p.Next = q
	}
	return root
}
