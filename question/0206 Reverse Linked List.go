package question

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	NN := &ListNode{
		Val:  -1,
		Next: nil,
	}
	p := head
	q := head
	for p != nil {
		q = p.Next
		p.Next = NN.Next
		NN.Next = p
		p = q
	}
	return NN.Next
}
