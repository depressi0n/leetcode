package question

// 给定一个二叉树,填充它的每个 next 指针，让这个指针指向其下一个右侧节点。如果找不到下一个右侧节点，则将 next 指针设置为 NULL。
//初始状态下，所有 next 指针都被设置为 NULL。 //

//进阶：
//你只能使用常量级额外空间。
//使用递归解题也符合要求，本题中递归程序占用的栈空间不算做额外的空间复杂度。

type Node117 struct {
	Val   int
	Left  *Node117
	Right *Node117
	Next  *Node117
}

func connect117(root *Node117) *Node117 {
	return connect117Core1(root)
}

// 主要思想：同样是递归处理，首先对子树进行处理（递归至最底层），递归出口情况为节点是空或者是叶子节点，直接返回
// 递归处理左子树和右子树
// 随后处理每一层，这里和完美二叉树不同的是，右孩子中每层最左不一定在同一侧，比如右孩子左子树就一个节点，但右子树是一个右孩子链表，同样对于左孩子也是这样
// 基于这样的条件，问题转变为如何找到二叉树中左孩子的最右节点和右孩子的最左节点
// 对于左孩子的最右节点，我们寻找的方式是找到左孩子的最左节点，此时按照Next域可以找到当前层的最右节点
// 所以问题转化为如何寻找二叉树每层的最左节点（特殊情况仍旧是左孩子并没有右孩子深，此时最左节点出现在右孩子中）
// 一种合理的方式是根据最左节点一直往后判断，因为下一层一定通过上一层的孩子连接起来
func connect117Core1(root *Node117) *Node117 {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	connect117Core1(root.Left)
	connect117Core1(root.Right)
	if root.Left == nil || root.Right == nil {
		return root
	}
	// 下一层
	nextLeftMostLeft, nextRightMostLeft := root.Left, root.Right
	nextLeftMostLeft.Next = nextRightMostLeft
	for {
		// 寻找左孩子的最左节点
		p := nextLeftMostLeft
		for p != nextRightMostLeft && p.Left == nil && p.Right == nil {
			p = p.Next
		}
		// 因为之前已经连接起来了，所以这里不是判断nil，而是判断上一层的终点
		if p == nextRightMostLeft {
			break
		}
		if p.Left != nil {
			nextLeftMostLeft = p.Left
		} else if p.Right != nil {
			nextLeftMostLeft = p.Right
		}
		// 从最左节点找到最右节点
		p = nextLeftMostLeft
		for p.Next != nil {
			p = p.Next
		}
		// 寻找右孩子的最左节点
		q := nextRightMostLeft
		for q != nil && q.Left == nil && q.Right == nil {
			q = q.Next
		}
		if q == nil {
			break
		}
		if q.Left != nil {
			q = q.Left
		} else {
			q = q.Right
		}
		nextRightMostLeft = q
		p.Next = nextRightMostLeft
	}
	return root
}

func connect117Core2(root *Node117) *Node117 {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	if root.Left != nil {
		connect117Core2(root.Left)
	}
	if root.Right != nil {
		connect117Core2(root.Right)
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

// 借助上一层形成的指针，并且不再需要递归处理
func connect117Core3(root *Node117) *Node117 {
	if root == nil {
		return root
	}
	start := root
	for start != nil {
		var last,nextStart *Node117
		for p:=start;p!=nil;p=p.Next{
			if p.Left!=nil{
				if nextStart==nil{
					nextStart=p.Left
				}
				if last !=nil{
					last.Next=p.Left
				}
				last=p.Left
			}
			if p.Right!=nil{
				if nextStart==nil{
					nextStart=p.Right
				}
				if last !=nil{
					last.Next=p.Right
				}
				last=p.Right
			}
		}
		start=nextStart
	}
	return root
}
