package question

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func getMinimumDifference(root *TreeNode) int {
	min := math.MaxInt32
	if root.Left != nil {
		l := root.Left
		for l.Right != nil { //左边的最大值
			l = l.Right
		}
		if root.Val-l.Val < min {
			min = root.Val - l.Val
		}
		minInLeft := getMinimumDifference(root.Left)
		if minInLeft < min {
			min = minInLeft
		}
	}
	if root.Right != nil {
		l := root.Right
		for l.Left != nil { //右边的最小值
			l = l.Left
		}
		if root.Val-l.Val < min {
			min = root.Val - l.Val
		}
		minInRight := getMinimumDifference(root.Right)
		if minInRight < min {
			min = minInRight
		}
	}
	return min
}
