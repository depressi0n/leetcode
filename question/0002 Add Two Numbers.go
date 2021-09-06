package question

// 题目描述：给你两个 非空 的链表，表示两个非负的整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
//请你将两个数相加，并以相同形式返回一个表示和的链表。
//你可以假设除了数字 0 之外，这两个数都不会以 0开头。


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return addTwoNumbersCore(l1,l2)
}
// 借助carry存储进位，不要忘记处理最后一个进位
func addTwoNumbersCore(l1 *ListNode, l2 *ListNode) *ListNode {
	p,q:=l1,l2
	carry:=0
	head:=&ListNode{}
	cur:=head
	for p!=nil && q!=nil{
		cur.Next=&ListNode{
			Val:(p.Val+q.Val+carry)%10,
			Next:nil,
		}
		cur=cur.Next
		carry=(p.Val+q.Val+carry)/10
		p=p.Next
		q=q.Next
	}
	if q!=nil{
		p=q
	}
	for p!=nil {
		cur.Next=&ListNode{
			Val:(p.Val+carry)%10,
			Next:nil,
		}
		cur=cur.Next
		carry=(p.Val+carry)/10
		p=p.Next
	}
	// 处理最后一个进位
	if carry!=0{
		cur.Next=&ListNode{
			Val:carry,
			Next:nil,
		}
	}
	return head.Next
}
// TODO：优化思路，借助输入的空间，而不是重新申请空间，使用一个pool存储之前放下来的节点，可以利用链表+头插法实现