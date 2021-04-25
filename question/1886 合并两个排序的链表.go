package question

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	p := l1
	q := l2
	head := &ListNode{
		Val:  -1,
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
	}
	for p != nil {
		cur.Next = p
		p = p.Next
		cur = cur.Next
	}
	for q != nil {
		cur.Next = q
		q = q.Next
		cur = cur.Next
	}
	return head.Next
}
