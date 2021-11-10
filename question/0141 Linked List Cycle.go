package question

// 给定一个链表，判断链表中是否有环。
// 如果链表中有某个节点，可以通过连续跟踪 next 指针再次到达，则链表中存在环。 为了表示给定链表中的环，
// 我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意：pos 不作为参数进行传递，仅仅是为了标识链表的实际情况。
// 如果链表中存在环，则返回 true 。 否则，返回 false 。
// 你能用 O(1)（即，常量）内存解决此问题吗
func hasCycle(head *ListNode) bool {
	return hasCycleCore(head)
}

// 朴素思想：将所有的指针放入一个map中，如果遇到访问过的map，则说明有环，否则说明无环
// 优化思想：快慢指针，快指针一次走两下，慢指针一次走一下，如果无环，那么快指针会遇到nil，如果有环，快指针会遇到慢指针
// 快慢指针相遇原理如下：假设链表中节点个数为a，环中节点个数为r，剩下节点个数为b，有a=b+r成立
// 另外，由于快慢指针速度的关系，在环中必然会出现相遇的情况。
// 假设相遇的位置离进入环节点的距离为c，那么有b+r+c=2*(b+c) => r = b
func hasCycleCore(head *ListNode) bool {
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
