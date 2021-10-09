package question

// 给定一棵树的前序遍历 preorder 与中序遍历  inorder。请构造二叉树并返回其根节点。

func buildTree105(preorder []int, inorder []int) *TreeNode {
	return buildTree105Core2(preorder, inorder)
}
func buildTree105Core1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 { //终止条件
		return nil
	}
	// 找到root
	i := 0
	for ; i < len(inorder) && inorder[i] != preorder[0]; i++ {

	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	root.Left = buildTree105Core1(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree105Core1(preorder[i+1:], inorder[i+1:])
	return root
}
// TODO: 未完全理解
// 前序遍历中任意两个相邻的节点u，v，只有两种可能：（1）v是u的左孩子；（2）u没有左孩子，并且v是u的某个祖先节点（或u本身）的右孩子
// 对于第二种情况，要找到v的父节点，因为是右孩子，在中序遍历中出现的位置应该在其父节点之后，而前序遍历中必然出现在父节点之前
func buildTree105Core2(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}
	// 栈s，初始化存储根节点
	s := make([]*TreeNode, 0, 100)
	s = append(s, root)
	// 辅助指针，指向中序遍历的第一个节点
	index := 0
	// 依次枚举前序遍历中除第一个节点外的每个节点，如果恰好指向栈顶节点，则不断弹出栈顶节点并右移，并将当前节点作为最后一个弹出的节点的右孩子，否则将当前节点作为栈顶节点的左孩子
	// 每次枚举都将当前节点入栈
	for i := 1; i < len(preorder); i++ {
		node := s[len(s)-1]
		if node.Val != inorder[index] {
			node.Left = &TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			s = append(s, node.Left)
		} else {
			for len(s) != 0 && s[len(s)-1].Val == inorder[index] {
				node=s[len(s)-1]
				s=s[:len(s)-1]
				index++
			}
			node.Right=&TreeNode{
				Val:   preorder[i],
				Left:  nil,
				Right: nil,
			}
			s=append(s,node.Right)
		}
	}
	return root
}
