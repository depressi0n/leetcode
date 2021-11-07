package question

// 请编写一个函数，用于 删除单链表中某个特定节点 。在设计函数时需要注意，你无法访问链表的头节点 head ，只能直接访问 要被删除的节点 。
//题目数据保证需要删除的节点 不是末尾节点 。

func deleteNode(node *ListNode) {
	deleteNodeCore(node)
}

// 没有任何办法能找到之前的节点
// 所以只能通过修改数值的方式完成，只需要修改下一个节点的数值即可，然后删除下一个节点
func deleteNodeCore(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
