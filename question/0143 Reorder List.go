package question

import (
	"fmt"
)

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	p := head.Next
	head.Next = nil
	//先摘再合
	//计算长度
	q := p
	cnt := 0
	for q != nil {
		cnt++
		q = q.Next
	}
	q = p
	cnt = (cnt - 1) / 2
	for cnt != 0 {
		q = q.Next
		cnt--
	}

	r := q.Next
	q.Next = nil
	q = r
	// 对q进行逆置
	r = q.Next
	q.Next = nil
	for r != nil {
		t := r.Next
		r.Next = q
		q = r
		r = t
	}

	tmp := p
	for tmp != nil {
		fmt.Printf("%d->", tmp.Val)
	}
	fmt.Println()
	tmp = q
	for tmp != nil {
		fmt.Printf("%d->", tmp.Val)
	}
	fmt.Println()

	//合并
	r = head
	for q != nil {
		r.Next = p
		p = p.Next
		r = r.Next
		r.Next = q
		q = q.Next
		r = r.Next
	}
	return
}
