package question

// 存在一个按升序排列的链表，给你这个链表的头节点 head ，请你删除所有重复的元素，使每个元素 只出现一次 。
//返回同样按升序排列的结果链表。

func deleteDuplicates0083(head *ListNode) *ListNode {
	return deleteDuplicates0083Core(head)
}

// 同上一题目
func deleteDuplicates0083Core(head *ListNode) *ListNode {
	res := &ListNode{}
	tail := res
	p := head
	for p != nil {
		for tail != res && p != nil && tail.Val == p.Val {
			// 说明这个值已经添加过了
			p = p.Next
		}
		// 将p摘下来
		if p != nil {
			tail.Next = p
			p = p.Next
			tail = tail.Next
			tail.Next = nil
		}
	}
	return res.Next
}
