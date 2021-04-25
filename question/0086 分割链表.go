package question

func partition(head *ListNode, x int) *ListNode {
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
