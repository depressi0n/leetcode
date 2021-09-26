package question

// 给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	return reverseBetweenCore(head, m, n)
}
func reverseBetweenCore(head *ListNode, m int, n int) *ListNode {
	if head == nil {
		return head
	}
	//用尾插法和头插法
	//首先前面用尾插法，中间一段用头插法，最后再用尾查法
	//哨兵节点
	res := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := res
	q := head
	r := head.Next
	for i := 1; i < m; i++ { //前m个用尾插法
		p.Next = q
		p = p.Next
		p.Next = nil

		q = r
		r = r.Next
	}
	//记住到来的尾部节点
	tail := q
	//中间这段用头插法，但要不断链
	//当前处理的就是q
	for i := 0; i < n-m+1; i++ {
		q.Next = p.Next
		p.Next = q

		q = r
		if r != nil {
			r = r.Next
		}
	}
	tail.Next = q
	return res.Next
}
