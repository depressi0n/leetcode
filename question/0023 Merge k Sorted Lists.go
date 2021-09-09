package question

import (
	"container/heap"
)

// 题目描述：给你一个链表数组，每个链表都已经按升序排列。
//请你将所有链表合并到一个升序链表中，返回合并后的链表。

func mergeKLists(lists []*ListNode) *ListNode {
	return mergeKListsCore4(lists)
}

// huffman树的形式
func mergeKListsCore1(lists []*ListNode) *ListNode {
	if len(lists) == 0 || lists[0] == nil {
		return nil
	}
	res := lists[0]
	for i := 1; i < len(lists); i++ {
		res = mergeTwoListsCore(res, lists[i])
	}
	return res
}

// 归并树的形式
func mergeKListsCore2(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}
	return mergeTwoListsCore(mergeKLists(lists[:len(lists)/2]), mergeKLists(lists[len(lists)/2:]))
}

type hSortList []*ListNode

func (h hSortList) Len() int {
	return len(h)
}

func (h hSortList) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}

func (h hSortList) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hSortList) Push(x interface{}) {
	*h=append(*h,x.(*ListNode))
}

func (h *hSortList) Pop() interface{} {
	old:=*h
	n:=len(old)
	x:=old[:n-1]
	*h=old[0:n-1]
	return x
}

// 优先级队列
func mergeKListsCore3(lists []*ListNode) *ListNode {
	h:=hSortList(lists)
	dummyHead:=new(ListNode)
	pre:=dummyHead
	for h.Len()>0{
		t:=heap.Pop(&h).(*ListNode) // TODO 判定有点问题，语言机制？
		if t.Next!=nil{
			heap.Push(&h,t.Next)
		}
		pre.Next=t
		pre=pre.Next
	}
	return dummyHead.Next

}

func mergeKListsCore4(lists []*ListNode) *ListNode {
	que:=make([]*ListNode,0)
	for i:=0;i<len(lists);i++{
		que=proQuePush(que,lists[i])
	}
	if len(que)<=0{
		return nil
	}
	start:=que[0]
	p:=start
	for {
		//fmt.Print(p.Val)
		que=proQuePush(que[1:],que[0].Next)
		if len(que)<=0{
			break
		}
		p.Next=que[0]
		p=p.Next
	}
	return start
}

func proQuePush(list []*ListNode,p *ListNode)[]*ListNode{
	if p==nil{
		return list
	}
	var minIndex int
	for minIndex=len(list)-1;minIndex>=0;minIndex--{
		if list[minIndex].Val<=p.Val{
			break
		}
	}
	if minIndex<0{
		tmp:=make([]*ListNode,len(list)+1)
		tmp[0]=p
		copy(tmp[1:],list)
		// for i:=0;i<len(tmp);i++{
		//     fmt.Print(tmp[i].Val," ")
		// }
		// fmt.Print(":",minIndex," ")
		return tmp
	}else{
		tmp:=make([]*ListNode,len(list)+1)
		copy(tmp[0:minIndex+1],list[0:minIndex+1])
		tmp[minIndex+1]=p
		copy(tmp[minIndex+2:],list[minIndex+1:])

		// for i:=0;i<len(tmp);i++{
		//     fmt.Print(tmp[i].Val," ")
		// }
		// fmt.Print(":",minIndex," ")
		return tmp
	}

}