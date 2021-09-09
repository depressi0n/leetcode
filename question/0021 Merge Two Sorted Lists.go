package question
//

func mergeTwoLists0021(l1 *ListNode, l2 *ListNode) *ListNode {
	return mergeTwoListsCore(l1,l2)
}
// 一次合并一个
func mergeTwoListsCore(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	res:=&ListNode{
		Val:  -1,
		Next: nil,
	}
	p,q:=l1,l2
	r:=res
	for p!=nil && q!=nil{
		if p.Val<q.Val{
			r.Next=p
			p=p.Next

		}else{
			r.Next=q
			q=q.Next
		}
		r=r.Next
		r.Next=nil
	}
	if p!=nil{
		r.Next=p
	}
	if q!=nil{
		r.Next=q
	}
	return res.Next
}