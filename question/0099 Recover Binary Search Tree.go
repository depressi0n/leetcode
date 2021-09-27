package question

// 给你二叉搜索树的根节点 root ，该树中的两个节点被错误地交换。请在不改变其结构的情况下，恢复这棵树。
func recoverTree(root *TreeNode) {
	recoverTreeCore2(root)
}

// 主要思想：二叉搜索树的中序遍历是有序的，所以如果有两个节点的值错误的交换，必然在序列中体现
func recoverTreeCore1(root *TreeNode) {
	//进行中序遍历
	// 用max存放当前遇到的最大值
	// 用bad存放可能坏值
	var max, bad, may *TreeNode
	p := root
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//check
			if max == nil || p.Val >= max.Val { //第一次和正常更新
				max = p
			} else { //异常情况发生
				if bad == nil { //第一次出问题,切换max，然后记录bad
					bad = max
					may = p
					max = p
				} else { //第二次出问题的时候做交换
					bad.Val, p.Val = p.Val, bad.Val
					return
				}
			}
			p = p.Right
		}
	}
	// 发生了相邻交换，所以只会进去一次
	bad.Val, may.Val = may.Val, bad.Val
	return
}

//进阶：使用 O(n) 空间复杂度的解法很容易实现。你能想出一个只使用常数空间的解决方案吗？
//TODO:morris 遍历
func recoverTreeCore2(root *TreeNode) {
	p := root
	var x, y, pred *TreeNode
	for p != nil {
		if p.Left != nil {
			mostRight := p.Left
			// 到最右边
			for mostRight.Right != nil && mostRight.Right != p {
				mostRight = mostRight.Right
			}
			if mostRight.Right == p {
				// 说明左子树访问完了
				// visit
				if pred != nil && p.Val < pred.Val {
					// 说明这里出现了点错误
					y = p
					if x == nil {
						x = pred
					}
				}
				// 转向右
				// 断开连接
				pred = p
				mostRight.Right = nil
				p = p.Right
			} else {
				// 连接，并访问左子树
				mostRight.Right = p
				p = p.Left
			}
		} else {
			// 没有左孩子，则访问根节点，转向右孩子
			// visit
			if pred != nil && p.Val < pred.Val {
				// 说明这里出现了点错误
				y = p
				if x == nil {
					x = pred
				}
			}
			pred = p
			p = p.Right
		}
	}
	x.Val, y.Val = y.Val, x.Val
}
