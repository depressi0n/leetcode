package question

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// TODO:可以弄一个pool，减少内存分配
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	nums := make([]int, n+1)
	for i := 1; i <= n; i++ {
		nums[i] = i
	}
	var generateSubTree func(int, int) []*TreeNode
	generateSubTree = func(left, right int) []*TreeNode {
		if left > right {
			return []*TreeNode{nil}
		}
		res := make([]*TreeNode, 0)
		for i := left; i <= right; i++ {
			leftTrees := generateSubTree(left, i-1)
			rightTrees := generateSubTree(i+1, right)
			for j := 0; j < len(leftTrees); j++ {
				for k := 0; k < len(rightTrees); k++ {
					root := &TreeNode{
						Val:   nums[i],
						Left:  nil,
						Right: nil,
					}
					root.Left = leftTrees[j]
					root.Right = rightTrees[k]
					res = append(res, root)
				}
			}
		}
		return res
	}
	//递归选择root
	res := generateSubTree(1, n)
	return res
}
