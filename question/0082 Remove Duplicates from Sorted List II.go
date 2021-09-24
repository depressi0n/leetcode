package question

//存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除链表中所有存在数字重复情况的节点，只保留原始链表中 没有重复出现 的数字。
//返回同样按升序排列的结果链表。
func deleteDuplicates(head *ListNode) *ListNode {
	return deleteDuplicatesCore2(head)
}
func deleteDuplicatesCore1(head *ListNode) *ListNode {
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
// 使用头节点
func deleteDuplicatesCore2(head *ListNode) *ListNode {
	res := &ListNode{}
	tail := res
	p := head
	for p != nil {
		if p.Next != nil && p.Next.Val == p.Val {
			r := p.Next
			for r != nil && r.Val == p.Val {
				r = r.Next
			}
			p = r
			continue
		}
		tail.Next = p
		if p != nil  {
			p = p.Next
		}
		tail = tail.Next
		tail.Next = nil
	}
	return res.Next
}
