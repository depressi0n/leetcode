package question

func flatten(root *TreeNode) {
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
		p.Right = q.Right
		q.Right = q.Left
		q.Left = nil
		p = q.Right
	}
}
