package question

import (
	"fmt"
	"math"
)

func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil { //0,1
		return head
	}
	if head.Next.Next == nil { //2
		if head.Val < head.Next.Val {
			return head
		}
		p := head.Next
		p.Next = head
		head.Next = nil
		return p
	}
	//归并排序
	//双指针
	head = &ListNode{
		Val:  math.MinInt32,
		Next: head,
	}
	q, s := head, head
	for q.Next != nil && q.Next.Next != nil {
		q = q.Next.Next //一次走两步
		s = s.Next      //一次走一步
	}
	q = s.Next
	s.Next = nil
	sortList(head)
	for x := head; x.Next != nil; x = x.Next {
		fmt.Printf("%d->", x.Val)
		fmt.Println()
	}
	sortList(q)
	for x := q; x.Next != nil; x = x.Next {
		fmt.Printf("%d->", x.Val)
		fmt.Println()
	}
	mergeList2(head, q)
	return head.Next
}

func mergeList2(p *ListNode, q *ListNode) *ListNode {
	head := &ListNode{
		Val:  math.MinInt32,
		Next: nil,
	}
	cur := head
	for p != nil && q != nil {
		if p.Val < q.Val {
			cur.Next = p
			p = p.Next
		} else {
			cur.Next = q
			q = q.Next
		}
		cur = cur.Next
		cur.Next = nil
	}
	if p != nil {
		cur.Next = p
	}
	if q != nil {
		cur.Next = q
	}
	return head.Next
}
