package question

// 根据先序和中序进行遍历
func buildTree105(preorder []int, inorder []int) *TreeNode {
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
	root.Left = buildTree105(preorder[1:i+1], inorder[0:i])
	root.Right = buildTree105(preorder[i+1:], inorder[i+1:])
	return root
}

//TODO 迭代的思想 利用的是前序遍历手动栈的结果
