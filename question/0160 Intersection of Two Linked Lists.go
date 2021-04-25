package question

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	p := headA
	for p != nil {
		p = p.Next
		lenA++
	}
	lenB := 0
	q := headB
	for q != nil {
		q = q.Next
		lenB++
	}
	p = headA
	q = headB
	delta := 0
	if lenA > lenB {
		delta = lenA - lenB
		for i := 0; i < delta; i++ {
			p = p.Next
		}

	} else {
		delta = lenB - lenA
		for i := 0; i < delta; i++ {
			q = q.Next
		}
	}
	for p != nil && q != nil {
		if p == q {
			return p
		}
		p = p.Next
		q = q.Next
	}
	return nil
}

// 浪漫的解法
func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	p := headA
	q := headB
	for p != q {
		if p.Next == nil && q.Next == nil {
			return nil
		}
		if p.Next == nil {
			p = headB
		} else {
			p = p.Next
		}
		if q.Next == nil {
			q = headA
		} else {
			q = q.Next
		}
	}
	return p
}
