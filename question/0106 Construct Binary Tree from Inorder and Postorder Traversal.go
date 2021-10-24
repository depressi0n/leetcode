package question

// 根据一棵树的中序遍历与后序遍历构造二叉树。你可以假设树中没有重复的元素。
func buildTree106(inorder []int, postorder []int) *TreeNode {
	return buildTree106Core(inorder,postorder)
}
func buildTree106Core(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) { //应当进入错误处理
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}
	// 找到在中序根节点的位置
	index := 0
	for ; index < len(inorder) && inorder[index] != postorder[len(postorder)-1]; index++ {

	}
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree106Core(inorder[:index], postorder[:index]),
		Right: buildTree106Core(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}

// TODO 迭代的思想，将后序、中序反向，就变成反先序和反中序，就可以利用之前题目的思想来做
