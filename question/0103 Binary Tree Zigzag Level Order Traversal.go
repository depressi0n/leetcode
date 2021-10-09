package question

// 给定一个二叉树，返回其节点值的锯齿形层序遍历。（即先从左往右，再从右往左进行下一层遍历，以此类推，层与层之间交替进行）。

func zigzagLevelOrder(root *TreeNode) [][]int {
	return zigzagLevelOrderCore(root)
}
func zigzagLevelOrderCore(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0, 100)
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	cur := 0
	start, end := 0, 1
	for cur < len(queue) {
		if queue[cur].Left != nil {
			queue = append(queue, queue[cur].Left)
		}
		if queue[cur].Right != nil {
			queue = append(queue, queue[cur].Right)
		}

		if cur == end-1 {
			tmp := make([]int, 0)
			if len(res)&1 == 0 {
				for ; start < end; start++ {
					tmp = append(tmp, queue[start].Val)
				}
			} else {
				for i := end - 1; i >= start; i-- {
					tmp = append(tmp, queue[i].Val)
				}
			}

			res = append(res, tmp)
			start = end
			end = len(queue)
		}
		cur++
	}
	return res
}
