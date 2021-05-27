package question

// TODO：这种结构的表示...似乎有问题
// TODO: 果然是线段树...
//type node struct {
//	left   int
//	right int
//	height int
//	next   *node
//}
//
//func (list *node) Insert(left int, height int) int {
//	p := list
//	if p.next != nil && p.next.right <= left {
//		p = p.next
//	}
//	if p==list{
//		tmp:=&node{
//			left:   left,
//			right:  p.next.left,
//			height: p.next.height + height, //新的高度
//			next:   p.next,
//		}
//		p.next=tmp
//		tmp.next.left=left+height //原来的高度
//		return p.next.height
//	}
//	//
//	if p.next == nil || left+height < p.next.left {
//		// 插入一个新的
//		tmp := &node{
//			left:   left,
//			right: left + height,
//			height: height,
//			next:   p.next,
//		}
//		p.next = tmp
//		return height
//	}
//	if left+height<=p.next.right{
//		// 拆成三份?
//		var tmp1 *node
//		if left+height< p.next.right{
//			tmp1=&node{
//				left:   left + height,
//				right:  p.next.right,
//				height: p.next.height,
//				next:   p.next.next,
//			}
//		}else{
//			tmp1=p.next.next
//		}
//		tmp2:=&node{
//			left:   left,
//			right:  left + height,
//			height: p.next.height + height,
//			next:   tmp1,
//		}
//		var tmp3 *node
//		if left+height == p.next.left{
//			tmp3=&node{
//				left:   p.next.left,
//				right:  left,
//				height: p.next.height,
//				next:   tmp2,
//			}
//		}else{
//			tmp3=p
//			tmp3.right=left
//		}
//		p.next=tmp3
//		return p.next.height+height
//	}
//	if p.next.left <= left && left+height <= p.next.right {
//		p.next.height += height
//		return p.next.height
//	}
//	if p.next.left <= left+height && left+height <= p.next.right {
//		p.next.left=left
//		tmp := &node{
//			left:   left,
//			right: p.next.right,
//			height: p.next.height+height,
//			next:   p.next,
//		}
//		p.next = tmp
//		return p.next.height+height
//	}
//	if p.next.left <left {
//		tmp:=&node{
//			left:   p.next.left,
//			right:  left,
//			height: p.next.height,
//			next:   p.next,
//		}
//		p.next=tmp
//		p=p.next
//	}else{
//		p.next.left=left
//	}
//	// 后面的可能需要合并
//	tmpH:=p.next.height
//	q:=p.next
//	for q.next!=nil &&q.next.right <left+height{
//		tmpH=max(tmpH,q.next.height)
//		q=q.next
//	}
//	if q.next == nil {
//		p.next.right=left+height
//		p.next.next=q.next
//		p.next.height=tmpH+height
//		return p.next.height
//	}else if left +height<=q.next.left {
//		p.next.right=q.right
//		p.next.next=q.next
//		p.next.height=tmpH+height
//		return p.next.height
//	}else{
//		q.next.left=left+height
//		p.next.right=left+height
//		p.next.next=q.next
//		q.next=nil
//		p.next.height=tmpH+height
//		return p.next.height
//	}
//}

//func FallingSquares(positions [][]int) []int {
//	lefts := make([]int, len(positions))
//	rights := make([]int, len(positions))
//	heights := make([]int, len(positions))
//	// 类似俄罗斯方块的叠放
//	// 目的：找到最大高度
//	res := make([]int, 0, len(positions))
//	// 思路：维护一个区间数组，保证插入新区间后合理拆分，并维护区间的高度
//	// 用链表去维护似乎更合适
//	for i := 0; i < len(positions); i++ {
//		if len(lefts) == 0 {
//			lefts = append(lefts, positions[i][0])
//			rights = append(rights, positions[i][0]+positions[i][1])
//			heights = append(heights, positions[i][1])
//			res = append(res, positions[i][1])
//			continue
//		}
//		left := positions[i][0]
//		right := positions[i][0] + positions[i][1]
//		j := 0
//		for ; j < len(lefts) && left <= rights[j]; j++ {
//
//		}
//		if j == 0 {
//			if right <= lefts[0] {
//				// 新插入一个在最前面
//				lefts = append([]int{left}, lefts...)
//				rights = append([]int{right}, rights...)
//				heights = append([]int{positions[i][1]}, heights...)
//				res = append(res, max(res[len(res)-1], positions[i][1]))
//				continue
//			}
//			if right < rights[0] {
//				if left < lefts[0] {
//					// 拆分后插入一个在最前面
//					lefts[0] = right
//					lefts = append([]int{left}, lefts...)
//					rights = append([]int{right}, rights...)
//					heights = append([]int{heights[0] + positions[i][1]}, heights...)
//					res = append(res, max(res[len(res)-1], positions[i][1]))
//					continue
//				} else if left == lefts[0]{
//					// 拆分成两份
//					lefts[0]=right
//					lefts = append([]int{left}, lefts...)
//					rights = append([]int{right}, rights...)
//					heights = append([]int{heights[0] + positions[i][1]}, heights...)
//					res = append(res, max(res[len(res)-1], heights[0] +positions[i][1]))
//					continue
//				}else  { //  left > lefts[0]
//					// 三份
//					lefts = append(lefts, left)
//					rights = append(rights, right)
//					heights=append(heights,heights[0]+positions[i][1])
//
//					lefts = append(lefts, right)
//					rights = append(rights, rights[0])
//					heights=append(heights,heights[0])
//
//					rights[0] = left
//					res = append(res, max(res[len(res)-1], heights[0] +positions[i][1]))
//					continue
//				}
//			} else if right==rights[0] {
//
//			}
//		}
//
//		res = append(res, max(res[len(res)-1], got))
//	}
//	return res
//}

type node699 struct {
	l, r, maxR, h int
	left, right   *node699
}

func FallingSquares(positions [][]int) []int {
	res := make([]int, len(positions))
	root := new(node699)
	maxH := 0
	var query func(root *node699, l, r int) int
	query = func(root *node699, l, r int) int {
		if root == nil || l >= root.maxR {
			return 0
		}
		curH := 0
		if !(r <= root.l || root.r <= l) {
			curH = root.h
		}
		curH = max(curH, query(root.left, l, r))
		if r >= root.l {
			curH = max(curH, query(root.right, l, r))
		}
		return curH
	}
	var insert func(root *node699, l, r int, h int) *node699
	insert = func(root *node699, l, r int, h int) *node699 {
		if root == nil {
			return &node699{
				l:     l,
				r:     r,
				maxR:  r,
				h:     h,
				left:  nil,
				right: nil,
			}
		}
		if l <= root.l {
			root.left = insert(root.left, l, r, h)
		} else {
			root.right = insert(root.right, l, r, h)
		}
		root.maxR = max(r, root.maxR)
		return root
	}
	for i := 0; i < len(positions); i++ {
		l := positions[i][0]
		r := positions[i][0] + positions[i][1]
		e := positions[i][1]
		curH := query(root, l, r)
		root = insert(root, l, r, curH+e)
		maxH = max(maxH, curH+e)
		res = append(res, maxH)
	}
	return res
}
