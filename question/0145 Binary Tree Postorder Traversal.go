package question

func postorderTraversal(root *TreeNode) []int {
	return postorderTraversalCore4(root)
}

// 递归思路
func postorderTraversalCore1(root *TreeNode) []int {
	//递归
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res
}

// 递归转成非递归，手动维护栈，采用先右后左再根的方式，最后逆置
func postorderTraversalCore2(root *TreeNode) []int {
	//非递归
	if root == nil {
		return []int{}
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

// 非递归的另一种方式，采用栈+一个辅助指针prev，辅助指针维护的是之前访问过的节点，旨在不会出现重复访问的情况
// 在访问某一个节点之前，当前节点的右子树是prev或是nil，说明以当前节点为根的子树已经访问完毕，访问工作指针，更新prev，并将工作指针置空
func postorderTraversalCore3(root *TreeNode) []int {
	res := make([]int, 0, 100)
	s := make([]*TreeNode, 0, 100)
	var pre *TreeNode
	for len(s) != 0 || root != nil {
		for root != nil {
			s = append(s, root)
			root = root.Left
			continue
		}
		root = s[len(s)-1]
		s = s[:len(s)-1]
		if root.Right == nil || root.Right == pre {
			res = append(res, root.Val)
			pre = root
			root = nil
		} else {
			s = append(s, root)
			root = root.Right
		}

	}
	return res
}

// Morris遍历
// 建立"连接"的过程无法省去，按照"右枝条"的方式去逆序完成结果的打印
func postorderTraversalCore4(root *TreeNode) (res []int) {
	if root == nil {
		return []int{}
	}
	reverse := func(a []int) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}
	// 添加路径的时候采用的方式是从当前节点的左子树一直往右走，而不是从当前节点
	addPath := func(node *TreeNode) {
		resSize := len(res)
		for ; node != nil; node = node.Right {
			res = append(res, node.Val)
		}
		reverse(res[resSize:])
	}

	p := root
	for p != nil {
		// 如果左子树不为空，说明需要建立连接
		if mostRight := p.Left; mostRight != nil {
			for mostRight.Right != nil && mostRight.Right != p {
				mostRight = mostRight.Right
			}
			// 如果右子树为空，则说明尚未建立"连接"
			if mostRight.Right == nil {
				mostRight.Right = p
				p = p.Left
				continue
			}
			// 否则说明，左子树已经访问完了，删除"连接"
			mostRight.Right = nil
			// 注意是当前节点的左孩子一直往右走，而不是当前节点
			// 这是因为右子树甚至都没有建立连接，这样会导致根节点先于右子树打印
			addPath(p.Left)
		}
		p = p.Right
	}
	// 最后才把根节点加进去
	addPath(root)
	return
}
