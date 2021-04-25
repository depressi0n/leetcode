package question

func detectCycle(head *ListNode) *ListNode {
	// 快指针一次走两个
	// 慢指针一次走一个
	// 最后肯定会碰上，相当于后面追赶
	// 如果有环的话
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
