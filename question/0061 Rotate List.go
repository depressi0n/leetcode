package question

// 给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。

func rotateRight(head *ListNode, k int) *ListNode {
	return rotateRightCore(head,k)
}

// 快慢指针的思想，找到新的head
// 首先快指针走k次，然后同步走到快指针指向结尾，最后将head指向慢指针的后一个，快指针和head形成圈
func rotateRightCore(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	length := 1
	r := head
	for r.Next != nil {
		r = r.Next
		length++
	}
	k = k % length
	if k == 0 {
		return head
	}
	r = head
	for i := 0; i < k; i++ {
		r = r.Next
	}
	q := head
	for r.Next != nil {
		r = r.Next
		q = q.Next
	}
	p := q.Next
	q.Next = r.Next
	r.Next = head
	head = p
	return head
}
