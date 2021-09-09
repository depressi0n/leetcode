package question

import "fmt"

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	return removeNthFromEndCore(head,n)
}
func removeNthFromEndCore(head *ListNode, n int) *ListNode {
	head=&ListNode{
		Val:  -1,
		Next: head,
	}
	p:=head
	// 先走n次
	for i:=0;i<n;i++{
		if p!=nil{
			p=p.Next
		}else{
			fmt.Println("error")
			return nil
		}
	}
	q:=head
	for p.Next!=nil{
		p=p.Next
		q=q.Next
	}
	q.Next=q.Next.Next
	return head.Next
}