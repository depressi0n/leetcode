package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func IsCousins(root *TreeNode, x int, y int) bool {
	// 堂兄弟，关键点在于还需要知道父亲是谁...
	// 层次遍历
	queue := make([]*TreeNode, 1, 100)
	queue[0] = root
	start := 0
	end := 1
	cur := 0
	flag := false
	var firstMeetNodeFather *TreeNode = nil
	for start < end {
		// 判断
		if queue[cur].Left != nil {
			queue = append(queue, queue[cur].Left)
			if queue[cur].Left.Val == x || queue[cur].Left.Val == y {
				if !flag {
					flag = true
					firstMeetNodeFather = queue[cur]
				} else {
					if firstMeetNodeFather == queue[cur] {
						return false
					} else {
						return true
					}
				}
			}
		}
		if queue[cur].Right != nil {
			queue = append(queue, queue[cur].Right)
			if queue[cur].Right.Val == x || queue[cur].Right.Val == y {
				if !flag {
					flag = true
					firstMeetNodeFather = queue[cur]
				} else {
					if firstMeetNodeFather == queue[cur] {
						return false
					} else {
						return true
					}
				}
			}
		}
		cur++
		if cur == end {
			// 切换到下一层之前已经发现了一个，但之前没有返回，说明不在同一层上，可以直接返回
			if flag {
				return false
			}
			start = end
			end = len(queue)
		}
	}
	// 可能到最后也没有这个两个节点
	return false
}
