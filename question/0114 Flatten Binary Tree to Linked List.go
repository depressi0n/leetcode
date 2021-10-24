package question

// 给你二叉树的根结点 root ，请你将它展开为一个单链表：
//展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
//展开后的单链表应该与二叉树 先序遍历 顺序相同。

func flatten(root *TreeNode) {
	flattenCore(root)
}

// 先序遍历的结果其实就是将左子树先潜入到右子树，然后递归的处理即可
func flattenCore(root *TreeNode) {
	if root == nil {
		return
	}
	flattenCore(root.Left)
	flattenCore(root.Right)
	if root.Left != nil {
		p := root.Left
		for p.Right != nil {
			p = p.Right
		}
		p.Right = root.Right
		root.Right = root.Left
		root.Left = nil
	}
	return
}

// 另一种方式是非递归的，根据先序遍历的路径来构造结果--首先找到左子树最右节点，将其右子树（现在为nil）指向跟节点的右子树
// 本质上将上面的递归改成非递归
func flattenCore2(root *TreeNode) {
	if root == nil {
		return
	}
	q := root
	for q != nil {
		if q.Left == nil {
			q = q.Right
			continue
		}
		p := q.Left
		//找到左子树最右节点
		for p.Right != nil {
			p = p.Right
		}
		// 将左子树最右节点拆到右子树树上，这是因为右子树上节点顺序在先序遍历中正好在左子树最右节点之后
		p.Right = q.Right
		// 此时右子树已经没有了，于是将左子树换到右子树
		q.Right = q.Left
		// 左子树置空
		q.Left = nil
		// 向下一层
		p = q.Right
	}
}
