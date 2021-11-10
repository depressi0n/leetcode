package question

//给你二叉树的根节点 root ，返回它节点值的 前序 遍历。
func preorderTraversal0144(root *TreeNode) []int {
	return preorderTraversal0144Core4(root)
}

// 递归思路：根节点->左子树->右子树
func preorderTraversal0144Core1(root *TreeNode) []int {
	//递归
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	res = append(res, preorderTraversal0144Core1(root.Left)...)
	res = append(res, preorderTraversal0144Core1(root.Right)...)
	return res
}

// 递归转成非递归：手动维护栈
func preorderTraversal0144Core2(root *TreeNode) []int {
	//非递归
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	s := make([]*TreeNode, 0)
	p := root
	for len(s) != 0 || p != nil {
		// 如果当前节点不为空，一直压栈到子树的最左节点，此时for循环结束
		for p != nil {
			s = append(s, p)
			res = append(res, p.Val)
			p = p.Left
		}
		// 在此处，左子树必为空，从栈中取出一个节点，转向其右子树
		p = s[len(s)-1]
		s = s[:len(s)-1]
		p = p.Right
	}
	return res
}

// 递归转成非递归的另一种写法，本质上与上面的写法一致
func preorderTraversal0144Core3(root *TreeNode) []int {
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

// Morris遍历：利用指向nil的节点去存储"栈"，在进入当前节点为根的左子树之前，将左子树和根"连接"起来，这样一直建立"连接"，直到左子树的最左，此时，左子树为空，需要转向右子树，此时的右子树中最右节点将和"根"节点有"连接"
// 借助建立好的"连接"，可以一直从最左侧以访问右子树的形式回到上层节点
// 只是在离开左子树之前，将"连接"断开
// 主要思想：一个工作指针p，借助mostRight存储当前节点为根的子树的最右节点，将其与根"连接"，在第二次遇到跟节点时即从左子树返回的时候，将"连接"断开
func preorderTraversal0144Core4(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	vals := make([]int, 0)
	var p, mostRight *TreeNode = root, nil
	// 以p为根的子树，进入for之前其实就是root为根的子树
	for p != nil {
		// 查看左子树是否为空
		mostRight = p.Left
		if mostRight != nil {
			// 如果左子树不为空，就将mostRight指向最右节点
			// 此时因为会有进入左子树和从左子树返回的情况
			// 如果是进入左子树，那么是需要建立连接的
			// 如果是从左子树返回，那么是需要删除连接的
			for mostRight.Right != nil && mostRight.Right != p {
				mostRight = mostRight.Right
			}
			// 如果此时mostRight是空，说明还没"连接"，建立"连接"
			// 前序遍历的时候，可以将根节点的值存入结果中，建立"连接"，进入左子树，继续建立"连接"
			if mostRight.Right == nil {
				vals = append(vals, p.Val)
				mostRight.Right = p
				p = p.Left
				continue
			}else{
				mostRight.Right = nil
			}
		} else {
			// 当左子树为空时，将根节点的值存入结果
			vals = append(vals, p.Val)
		}
		// 转向去遍历右子树
		p = p.Right
	}
	return vals
}
