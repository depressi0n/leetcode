package question

//从排序后的数组创建平衡二叉树，递归处理即可
func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{
			Val:   nums[0],
			Left:  nil,
			Right: nil,
		}
	}
	if len(nums) == 2 {
		return &TreeNode{
			Val: nums[1],
			Left: &TreeNode{
				Val:   nums[0],
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		}
	}
	return &TreeNode{
		Val:   nums[(len(nums)-1)/2],
		Left:  sortedArrayToBST(nums[0 : (len(nums)-1)/2]),
		Right: sortedArrayToBST(nums[(len(nums)-1)/2+1:]),
	}
}
