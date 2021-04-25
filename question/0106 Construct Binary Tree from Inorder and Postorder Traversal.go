package question

func buildTree106(inorder []int, postorder []int) *TreeNode {
	if len(inorder) != len(postorder) { //应当进入错误处理
		return nil
	}
	if len(inorder) == 0 {
		return nil
	}
	index := 0
	for ; index < len(inorder) && inorder[index] == postorder[len(postorder)-1]; index++ {

	}
	return &TreeNode{
		Val:   postorder[len(postorder)-1],
		Left:  buildTree106(inorder[:index], postorder[:index]),
		Right: buildTree106(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}

// TODO 迭代的思想，将后序、中序反向，就变成反先序和反中序，就可以利用之前题目的思想来做
