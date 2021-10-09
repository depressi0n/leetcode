package question

// 给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
func levelOrder(root *TreeNode) [][]int {
	return levelOrderCore(root)
}
func levelOrderCore(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := make([][]int, 0,100)
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
			for ; start < end; start++ {
				tmp = append(tmp, queue[start].Val)
			}
			res = append(res, tmp)
			end = len(queue)
		}
		cur++
	}
	return res
}
