package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal144_1(root *TreeNode) []int {
	//递归
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, preorderTraversal(root.Left)...)
	res = append(res, preorderTraversal(root.Right)...)
	return res
}

func preorderTraversal144_2(root *TreeNode) []int {
	//非递归
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	s := make([]*TreeNode, 0)

	p := root

	for len(s) != 0 || p != nil {
		for p != nil {
			s = append(s, p)
			res = append(res, p.Val)
			p = p.Left
		}
		p = s[len(s)-1]
		s = s[:len(s)-1]
		p = p.Right
	}
	return res
}

func preorderTraversal144_3(root *TreeNode) []int {
	//另一种非递归
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	s := make([]*TreeNode, 0)

	p := root

	for len(s) != 0 || p != nil {
		if p != nil {
			s = append(s, p)
			res = append(res, p.Val)
			p = p.Left
			continue //本质上是一样的
		}
		p = s[len(s)-1]
		s = s[:len(s)-1]
		p = p.Right
	}
	return res
}

// Morris遍历 TODO：弄懂
func preorderTraversal(root *TreeNode) (vals []int) {
	var p1, p2 *TreeNode = root, nil
	for p1 != nil {
		p2 = p1.Left
		if p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				vals = append(vals, p1.Val)
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
		} else {
			vals = append(vals, p1.Val)
		}
		p1 = p1.Right
	}
	return
}
