package question

func swapPairs(head *ListNode) *ListNode {
	return swapPairsCore(head)
}
func swapPairsCore(head *ListNode) *ListNode {
	res:=&ListNode{
		Val:-1,
		Next:nil,
	}
	tail:=res
	for head!=nil{
		p:=head.Next
		tail.Next=head
		head.Next=nil
		if p==nil{
			break
		}
		head=p.Next
		p.Next=tail.Next
		tail.Next=p
		tail=tail.Next.Next
	}
	return res.Next
}