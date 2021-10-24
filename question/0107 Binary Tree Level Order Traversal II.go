package question

// 给定一个二叉树，返回其节点值自底向上的层序遍历。
// 即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历
func levelOrderBottom(root *TreeNode) [][]int {
	return levelOrderBottomCore(root)
}
func levelOrderBottomCore(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0, 100)
	queue := make([]*TreeNode, 1, 100)
	queue[0] = root
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
			tmp := make([]int, end-start)
			for i := 0; i < end-start; i++ {
				tmp[i] = queue[start+i].Val
			}
			if res == nil {
				res = append(res, tmp)
			} else {
				res = append([][]int{tmp}, res...)
			}
			start = end
			end = len(queue)
		}
		cur++
	}
	return res
}
