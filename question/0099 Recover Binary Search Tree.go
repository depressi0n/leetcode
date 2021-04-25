package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
//TODO:morris 遍历
func recoverTree(root *TreeNode) {
	//进行中序遍历
	var max, bad *TreeNode
	p := root
	stack := make([]*TreeNode, 0)
	for len(stack) != 0 || p != nil {
		if p != nil {
			stack = append(stack, p)
			p = p.Left
		} else {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			//check
			if max == nil || p.Val >= max.Val { //第一次和正常更新
				max = p
			} else { //异常情况发生
				if bad == nil { //第一次出问题,切换max，然后记录bad
					bad = max
					max = p
				} else { //第二次出问题的时候做交换
					bad.Val, p.Val = p.Val, bad.Val
					return
				}
			}
			p = p.Right
		}
	}
}
