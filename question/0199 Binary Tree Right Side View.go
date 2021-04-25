package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{root.Val}
	// bfs或者dfs都可以
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	end := root
	for len(queue) != 0 {
		if queue[0].Left != nil {
			queue = append(queue, queue[0].Left)
		}
		if queue[0].Right != nil {
			queue = append(queue, queue[0].Right)
		}
		if queue[0] == end {
			res = append(res, end.Val)
			end = queue[len(queue)-1]
		}
		queue = queue[1:]
	}
	return res
}

//TODO: dfs
