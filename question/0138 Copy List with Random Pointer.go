package question

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Next *Node
 *     Random *Node
 * }
 */
type Node138 struct {
	Val    int
	Next   *Node138
	Random *Node138
}

func copyRandomList(head *Node138) *Node138 {
	//x := make(map[*Node138]*Node138)
	//index := 0
	//p := head
	//for p != nil {
	//	x[p] = &Node138{
	//		Val:    p.Val,
	//		Next:   nil,
	//		Random: nil,
	//	}
	//	index++
	//	p = p.Next
	//}
	//p = head
	//for p != nil {
	//	if p.Next != nil {
	//		x[p].Next = x[p.Next]
	//	}
	//	if p.Random != nil {
	//		x[p].Random = x[p.Random]
	//	}
	//	p = p.Next
	//}
	//return x[head]

	//优化：把每个节点都复制一边
	//然后拆出来就可以
	if head == nil {
		return nil
	}
	for node := head; node != nil; node = node.Next.Next {
		node.Next = &Node138{
			Val:    node.Val,
			Next:   node.Next,
			Random: nil,
		}
	}
	for node := head; node != nil; node = node.Next.Next {
		if node.Random != nil {
			node.Next.Random = node.Random.Next
		}
	}
	// 拆出来
	ret := head.Next
	for node := head; node != nil; node = node.Next {
		t := node.Next
		node.Next = node.Next.Next
		if t.Next != nil {
			t.Next = t.Next.Next
		}
	}
	return ret
}
