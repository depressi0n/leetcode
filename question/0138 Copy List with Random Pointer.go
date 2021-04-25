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
	x := make(map[*Node138]*Node138)
	index := 0
	p := head
	for p != nil {
		x[p] = &Node138{
			Val:    p.Val,
			Next:   nil,
			Random: nil,
		}
		index++
		p = p.Next
	}
	p = head
	for p != nil {
		if p.Next != nil {
			x[p].Next = x[p.Next]
		}
		if p.Random != nil {
			x[p].Random = x[p.Random]
		}
		p = p.Next
	}
	return x[head]
}
