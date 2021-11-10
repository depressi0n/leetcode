package question

// 给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。注意，pos 仅仅是用于标识环的情况，并不会作为参数传递到函数中。
//说明：不允许修改给定的链表。
//你是否可以使用 O(1) 空间解决此题？
func detectCycle(head *ListNode) *ListNode {
	return detectCycleCore(head)
}

// 朴素思想：将所有的指针放入一个map中，如果遇到访问过的map，则说明有环，否则说明无环
// 优化思想：快慢指针，快指针一次走两下，慢指针一次走一下，如果无环，那么快指针会遇到nil，如果有环，快指针会遇到慢指针
// 快慢指针相遇原理如下：假设链表中节点个数为a，环中节点个数为r，剩下节点个数为b，有a=b+r成立
// 另外，由于快慢指针速度的关系，在环中必然会出现相遇的情况。
// 为什么慢指针在相遇时没有走完一圈：c % r == (x +2*c)%r. r*n+c=b+2c => c = b- nr ，当n=1时 0<=c<r，所以必然在第一圈相遇
// 假设相遇的位置离进入环节点的距离为c，那么有b+nr+c=2*(b+c) => nr = b+c => c= b - nr
func detectCycleCore(head *ListNode) *ListNode {
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
