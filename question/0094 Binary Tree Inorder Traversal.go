package question

// 给定一个二叉树的根节点 root ，返回它的 中序 遍历。
func inorderTraversal(root *TreeNode) []int {
	return inorderTraversalCore4(root)
}

func inorderTraversalCore1(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0, 100)
	res = append(res, inorderTraversalCore1(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversalCore1(root.Right)...)
	return res
}
func inorderTraversalCore2(root *TreeNode) []int {
	res := make([]int, 0, 100)
	s := make([]*TreeNode, 0, 100)
	p := root
	for len(s) != 0 || p != nil {
		for p != nil {
			s = append(s, p)
			p = p.Left
		}
		p = s[len(s)-1]
		s = s[:len(s)-1]
		res = append(res, p.Val)
		p = p.Right
	}
	return res
}
func inorderTraversalCore3(root *TreeNode) []int {
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

// Morris 遍历
// 如果当前节点没有左孩子，则将当前节点值加入结果，再访问其右孩子
// 如果当前节点有左孩子，则找到左子树树上最右节点，记为pre，根据pre右孩子是否为空，进行操作
//    如果pre的右孩子为空，则将其右孩子指向当前节点，然后访问当前节点的左孩子
//    如果pre的右孩子不为空，则此时指向当前节点，说明左子树已经遍历完，将pre的右孩子置空，将当前节点的值加入结果，访问当前节点的右孩子
func inorderTraversalCore4(root *TreeNode) []int {
	res := make([]int, 0, 100)
	p := root
	var mostRight *TreeNode
	for p != nil {
		if p.Left == nil {
			res = append(res, p.Val)
			p = p.Right
		} else {
			mostRight = p.Left
			for mostRight.Right != nil && mostRight.Right != p {
				mostRight = mostRight.Right
			}
			// 左子树已经遍历完
			if mostRight.Right == p {
				mostRight.Right = nil
				res=append(res,p.Val)
				p = p.Right
			} else {
				mostRight.Right = p
				p = p.Left
			}
		}
	}
	return res
}