package question

func postorderTraversalRec(root *TreeNode) []int {
	//递归
	if root == nil {
		return nil
	}
	res := make([]int, 0)
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}
func postorderTraversalNotRec(root *TreeNode) []int {
	//非递归
	if root == nil {
		return nil
	}
	preRes := make([]int, 0)
	res := make([]int, 0)

	s := make([]*TreeNode, 0)
	p := root
	for len(s) != 0 || p != nil {
		for p != nil {
			preRes = append(preRes, p.Val)
			s = append(s, p)
			p = p.Right
		}
		p = s[len(s)-1]
		s = s[:len(s)-1]
		p = p.Left
	}
	for i := len(preRes) - 1; i >= 0; i-- {
		res = append(res, preRes[i])
	}
	return res
}

//只用一个栈的递归
func postorderTraversal1(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev { //右孩子为空或者右孩子是prev，访问之
			res = append(res, root.Val)
			prev = root
			root = nil
		} else { //再一次直接放入栈中
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

//Morris遍历，TODO：弄懂
func postorderTraversal(root *TreeNode) (res []int) {
	reverse:=func (a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p1 := root
	for p1 != nil {
		if p2 := p1.Left; p2 != nil {
			for p2.Right != nil && p2.Right != p1 {
				p2 = p2.Right
			}
			if p2.Right == nil {
				p2.Right = p1
				p1 = p1.Left
				continue
			}
			p2.Right = nil
			addPath(p1.Left)
		}
		p1 = p1.Right
	}
	addPath(root)
	return
}



func postorder(root *TreeNode)[]int{
	res:=make([]int,0,100)
	s:=make([]*TreeNode,0,100)
	var pre *TreeNode
	for len(s)!=0 || root !=nil{
		for root!=nil{
			s=append(s,root)
			res=append(res,root.Val)
			root=root.Left
			continue
		}
		root=s[len(s)-1]
		s=s[:len(s)-1]
		if root.Right == nil || root.Right==pre{
			res=append(res,root.Val)
			pre=root
			root=nil
		} else{
			s=append(s,root)
			root=root.Right
		}

	}
	return res
}