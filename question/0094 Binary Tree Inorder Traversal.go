package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
	res := make([]int, 0)
	//树的中序遍历
	stack := make([]*TreeNode, 0)
	p := root
	for len(stack) != 0 || p != nil {
		if p == nil { //意味着栈肯定不空
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			res = append(res, p.Val)
			p = p.Right
		} else {
			stack = append(stack, p)
			p = p.Left
		}
	}
	return res
}

// 另一种 中序遍历
func inorderTraversal_1(root *TreeNode) []int {
	if root==nil{
		return nil
	}
	res:=make([]int,0,100)
	s:=make([]*TreeNode,0,100)
	for len(s)!=0 || root!=nil{
		for root!=nil{
			s=append(s,root)
			root=root.Left
		}
		root=s[len(s)-1]
		s=s[:len(s)-1]
		res=append(res,root.Val)
		root=root.Right
	}
	return res
}