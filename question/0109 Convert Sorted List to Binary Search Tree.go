package question

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func sortedListToBST(head *ListNode) *TreeNode {
	length := 0
	p := &ListNode{Next: head}
	for p.Next != nil { //如果head不是nil，长度肯定能正确算出来
		length++
		p = p.Next
	}
	var generateTreeFromList func(*ListNode, int) *TreeNode
	generateTreeFromList = func(head *ListNode, length int) *TreeNode {
		if length == 0 {
			return nil
		}
		if length == 1 {
			return &TreeNode{
				Val:   head.Val,
				Left:  nil,
				Right: nil,
			}
		}
		if length == 2 {
			return &TreeNode{
				Val: head.Next.Val,
				Left: &TreeNode{
					Val:   head.Val,
					Left:  nil,
					Right: nil,
				},
				Right: nil,
			}
		}
		q := head
		for i := 0; i < length/2; i++ {
			q = q.Next
		}
		return &TreeNode{
			Val:   q.Val,
			Left:  generateTreeFromList(head, length/2),
			Right: generateTreeFromList(q.Next, length-length/2-1),
		}
	}
	return generateTreeFromList(head, length)
}
