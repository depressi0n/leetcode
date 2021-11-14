package question

import (
	"math"
)

// 给你链表的头结点head，请将其按 升序 排列并返回 排序后的链表 。
//进阶：
//你可以在O(nlogn) 时间复杂度和常数级空间复杂度下，对链表进行排序吗？

func sortList(head *ListNode) *ListNode {
	return sortListCore2(head)
}

// 主要思想：时间复杂度O(nlogn)，空间复杂度O(1)
// 考虑在数组前提下，时间复杂度O(nlogn)的排序算法：快速排序，归并排序，堆排序
// 考虑快速排序，因为链表的缘故，必然需要使用递归的方式去处理左右两边的结果，不太可能
// 考虑归并排序，先处理两边，最后归并（Top to Down），也要借助递归的方式，但是因为归并排序的结果是一个尾递归，所以可以改成Down to Top的方式
// 考虑堆排序，这里不允许借助辅助空间，无法建立堆？暂时无法考虑

// 以下是递归的思路
func sortListCore(head *ListNode) *ListNode {
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
	sortListCore(head)
	//for x := head; x.Next != nil; x = x.Next {
	//	fmt.Printf("%d->", x.Val)
	//	fmt.Println()
	//}
	sortListCore(q)
	//for x := q; x.Next != nil; x = x.Next {
	//	fmt.Printf("%d->", x.Val)
	//	fmt.Println()
	//}
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

// 下面用非递归的方式实现
func sortListCore2(head *ListNode) *ListNode {
	// 首先求出链表的长度
	length := 0
	p := head
	for p != nil {
		p = p.Next
		length++
	}
	if length <= 1 {
		return head
	}
	// 初始步长为1，已经排序结束
	// 步长为2，每次循环后翻倍，如果翻倍后的长度比length大，即为最后一次
	for step := 2; step < length<<1; step <<= 1 {
		fakeHead := &ListNode{
			Val:  0,
			Next: nil,
		}
		// 已经归并完成的结果
		r := fakeHead
		// 未完成归并的链表
		post := head
		for post != nil {
			// 可能的情况有：
			// (1)前半段和后半段长度均为step>>1
			// (2)前半段为step>>1但后半段小于step>>1
			// (3)前半段小于step>>1没有后半段，需要单独处理
			// 前面半段
			p = post
			for i := 1; i < step>>1 && p.Next != nil; i++ {
				p = p.Next
			}
			// 如果是第三种情况，直接接在r后即可
			if p.Next == nil {
				r.Next = post
				break
			}
			// 后面半段
			q := p.Next
			for i := 1; i < step>>1 && q.Next != nil; i++ {
				q = q.Next
			}
			// 将工作区间与未归并部分分开
			p, q, post = post, p.Next, q.Next
			// 开始归并
			curp, curq := p, q
			for curp != q && curq != post {
				if curp.Val < curq.Val {
					r.Next = curp
					curp = curp.Next
					r = r.Next
					r.Next = nil
				} else {
					r.Next = curq
					curq = curq.Next
					r = r.Next
					r.Next = nil
				}
			}
			for curp != q {
				r.Next = curp
				curp = curp.Next
				r = r.Next
				r.Next = nil
			}
			for curq != post {
				r.Next = curq
				curq = curq.Next
				r = r.Next
				r.Next = nil
			}
		}
		head = fakeHead.Next
	}
	return head
}
