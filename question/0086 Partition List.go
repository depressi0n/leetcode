package question

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//你应当 保留 两个分区中每个节点的初始相对位置。
func partition(head *ListNode, x int) *ListNode {
	return partitionCore2(head, x)
}

// 基本思路，用一个tail串联记录比x值大或相等的的，比x小的直接串联在其后
func partitionCore1(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var less, more, lessEnd, moreEnd *ListNode
	insertInto := func(h, end *ListNode, cur *ListNode) (*ListNode, *ListNode) {
		if end == nil {
			end = cur
			h = end
			return h, end
		}
		end.Next = cur
		end = cur
		return h, end
	}
	p := head
	for p != nil {
		q := p
		p = p.Next
		q.Next = nil
		if q.Val < x {
			less, lessEnd = insertInto(less, lessEnd, q)
		} else {
			more, moreEnd = insertInto(more, moreEnd, q)
		}
	}
	if lessEnd == nil {
		return more
	}
	lessEnd.Next = more
	return less
}
func partitionCore2(head *ListNode, x int) *ListNode {
	tail := &ListNode{}
	r := tail
	head = &ListNode{
		Val:  0,
		Next: head,
	}
	p := head.Next
	pre := head
	for p != nil {
		if p.Val >= x {
			// 摘下来
			pre.Next = p.Next
			p.Next = nil

			r.Next = p
			r=r.Next

			p=pre.Next
		}else{
			pre=p
			p=p.Next
		}
	}
	// 最后接上
	pre.Next=tail.Next
	return head.Next
}
