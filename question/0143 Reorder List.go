package question

// 给定一个单链表 L 的头节点 head ，单链表 L 表示为：
// L0→ L1→ … → Ln-1→ Ln
// 请将其重新排列后变为：
// L0→Ln→L1→Ln-1→L2→Ln-2→ …
// 不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
func reorderList(head *ListNode) {
	reorderListCore(head)
}

// 主要思想：首先取得链表长度，然后取下后半段，头插法逆置，然后和前半段链表合并
func reorderListCore(head *ListNode) {
	// 特殊情况的判定，长度小于2直接返回即可
	if head == nil {
		return
	}
	p := head
	length := 1
	for p.Next != nil {
		p = p.Next
		length++
	}
	if length<=2{
		return
	}
	p = head
	// 如果长度为偶数，应该到第n/2个位置，否则应该到第(n+1)/2个
	for i := 1; i < (length+1)/2; i++ {
		p = p.Next
	}
	// 取下后半段，逆置
	post := p.Next
	p.Next = nil
	cur := post
	post = post.Next
	cur.Next = nil
	for cur != nil {
		cur.Next = p.Next
		p.Next = cur

		cur = post
		if cur != nil {
			post = post.Next
			cur.Next = nil
		}
	}
	// 合并
	post = p.Next
	p.Next = nil
	p = head
	for post != nil {
		cur = post
		post = post.Next
		cur.Next = nil

		cur.Next = p.Next
		p.Next = cur
		p = cur.Next
	}
	return
}
