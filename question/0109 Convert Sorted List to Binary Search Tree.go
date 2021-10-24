package question

// 给定一个单链表，其中的元素按升序排序，将其转换为高度平衡的二叉搜索树。
//本题中，一个高度平衡二叉树是指一个二叉树每个节点的左右两个子树的高度差的绝对值不超过 1。

func sortedListToBST(head *ListNode) *TreeNode {
	return sortedListToBSTCore(head)
}
// 使用递归的方式去构造
func sortedListToBSTCore(head *ListNode) *TreeNode {
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
