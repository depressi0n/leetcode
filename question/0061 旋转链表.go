package question

func rotateRight(head *ListNode, k int) *ListNode {
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
