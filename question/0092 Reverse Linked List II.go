package question

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
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
