package question

func deleteDuplicates(head *ListNode) *ListNode {
	var res, q *ListNode
	p := head
	for p != nil {
		if p.Next != nil && p.Next.Val == p.Val { //如果是重复值，则一直跳到不相同到一个
			r := p.Next.Next
			for r.Val == p.Val {
				r = r.Next
			}
			p = r
			continue
		}
		if res == nil {
			res = p
			q = p
		} else {
			q.Next = p
			q = q.Next
		}
		p = p.Next
		q.Next = nil
	}
	return res
}
