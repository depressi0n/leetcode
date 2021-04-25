package question

func hasCycle(head *ListNode) bool {
	// 快指针一次走两个
	// 慢指针一次走一个
	// 最后肯定会碰上，相当于后面追赶
	// 如果有环的话
	if head == nil {
		return false
	}
	quick, slow := head, head
	flag := false
	for quick.Next != nil && quick.Next.Next != nil && slow.Next != nil {
		quick = quick.Next.Next
		slow = slow.Next
		if quick == slow {
			flag = true
			break
		}
	}
	return flag
}
