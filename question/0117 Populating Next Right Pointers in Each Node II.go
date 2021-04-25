package question

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

//	TODO:层次遍历的思想，用next域作为层次的队列
func connect117(root *Node117) *Node117 {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil {
		connect117(root.Left)
	}
	if root.Right != nil {
		connect117(root.Right)
	}

	if root.Left == nil { //root.Left==nil
		return root
	}
	p := root.Left
	q := root.Right
	curLeft := root.Left
	flag := true
	for { //从第二层开始
		if p == nil || q == nil {
			return root
		}
		flag = true
		for curLeft != nil && curLeft != q && flag {
			flag = false
			if curLeft.Left != nil {
				curLeft = curLeft.Left
			} else if curLeft.Right != nil {
				curLeft = curLeft.Right
			} else {
				curLeft = curLeft.Next
				flag = true
			}
		}

		p.Next = q

		// 下一层左边找到最左节点
		if curLeft == nil {
			return root
		}
		p = curLeft
		for p.Next != nil {
			p = p.Next
		}

		// 下一层右边找到最左节点
		flag = true //必须进一轮
		for q != nil && flag {
			flag = false
			if q.Left != nil {
				q = q.Left
			} else if q.Right != nil {
				q = q.Right
			} else {
				q = q.Next
				flag = true
			}
		}
	}
	return root
}
