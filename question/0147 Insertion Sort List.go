package question

import (
	"math"
)

//另一条奇葩的思路是对值进行排序，然后填入链表中
func insertionSortList(head *ListNode) *ListNode {
	p := head
	// 虚拟头节点
	head = &ListNode{
		Val:  math.MinInt32,
		Next: nil,
	}
	for p != nil {
		q := p.Next //不断链
		p.Next = nil
		r := head
		for r.Next != nil && r.Next.Val < p.Val {
			r = r.Next
		}
		p.Next = r.Next
		r.Next = p
		p = q
	}
	return head.Next
}
