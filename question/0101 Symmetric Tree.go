package question
// 给定一个二叉树，检查它是否是镜像对称的。
func isSymmetric(root *TreeNode) bool {
	return isSymmetricCore2(root)
}

// 直接将root进行两次，一次往左，一次往右
func isSymmetricCore1(root *TreeNode) bool {
	return check(root, root)
}

//递归的思路
func check(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val && check(p.Left, q.Right) && check(p.Right, q.Left)
}

// 递推，根节点入队列两次，每次提取两个节点比较Val，如果先相同则前一个节点左右孩子入队，后一个节点右左入队，否则返回
// 入队时空节点也要入队，比较时
func isSymmetricCore2(root *TreeNode) bool {
	// 引入队列，初始化时根节点入队列两次，每次提取两个节点比较并将两个节点的左右孩子相反顺序入队
	q := make([]*TreeNode, 0, 100)
	q = append(q, root)
	q = append(q, root)
	for len(q) != 0 {
		u := q[0]
		v := q[1]
		q = q[2:]
		if u == nil && v == nil {
			continue
		}
		if (u == nil || v == nil) || (u.Val != v.Val) {
			return false
		}
		q = append(q, u.Left)
		q = append(q, v.Right)
		q = append(q, u.Right)
		q = append(q, v.Left)
	}
	return true
}
