package question

// 给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。
//高度平衡二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

//从排序后的数组创建平衡二叉树，递归处理即可
func sortedArrayToBST(nums []int) *TreeNode {
	return sortedArrayToBSTCore(nums)
}
func sortedArrayToBSTCore(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	return &TreeNode{
		Val:   nums[(len(nums)-1)/2],
		Left:  sortedArrayToBSTCore(nums[0 : (len(nums)-1)/2]),
		Right: sortedArrayToBSTCore(nums[(len(nums)-1)/2+1:]),
	}
}
