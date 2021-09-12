package question

// 题目描述：给你一个链表，每k个节点一组进行翻转，请你返回翻转后的链表。
//k是一个正整数，它的值小于或等于链表的长度。
//如果节点总数不是k的整数倍，那么请将最后剩余的节点保持原有顺序。
//进阶：
//你可以设计一个只使用常数额外空间的算法来解决此问题吗？
//你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换

func reverseKGroup(head *ListNode, k int) *ListNode {
	return reverseKGroupCore2(head, k)
}

// 递归的思路
func reverseKGroupCore1(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	p := head
	i := 0
	for ; i < k-1 && p != nil; i++ {
		p = p.Next
	}
	// 长度不够，直接返回
	if i < k-1 || p == nil {
		return head
	}
	// 截断
	remain := p.Next
	p.Next = nil
	// 通过头插入法反转这k个
	fakeHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	r := head
	var t *ListNode
	for r != nil {
		t = r.Next
		r.Next = fakeHead.Next
		fakeHead.Next = r
		r = t
	}
	// 连接后面的
	head.Next = reverseKGroupCore1(remain, k)
	return fakeHead.Next
}

// 既然是一个尾递归，不妨转成递推

func reverseKGroupCore2(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	fakeHead := &ListNode{
		Val:  -1,
		Next: nil,
	}
	tail:=fakeHead
	remain := head
	var p *ListNode
	for remain != nil {
		p=remain
		// 扫出k个
		i := 0
		for ; i < k-1 && p != nil; i++ {
			p = p.Next
		}
		// 长度不够，顺序插入之前的结果中
		if i < k-1 || p == nil {
			tail.Next=remain
			return fakeHead.Next
		}
		// 截断
		r:=remain
		remain = p.Next
		p.Next = nil
		// 通过对tail进行头插入法反转这k个
		for r != nil {
			p = r.Next
			r.Next = tail.Next
			tail.Next = r
			r = p
		}
		// tail移动到当前末尾
		for tail.Next!=nil{
			tail=tail.Next
		}
	}
	return fakeHead.Next
}
