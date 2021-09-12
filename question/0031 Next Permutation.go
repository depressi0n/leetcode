package question

// 题目描述:实现获取 下一个排列 的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列
//（即，组合出下一个更大的整数）。
//	如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。
//	必须 原地 修改，只允许使用额外常数空间。

func nextPermutation(nums []int) []int {
	return nextPermutationCore(nums)
}

// 这个主要是靠记住
// 从后往前找到第一个非递增（从后往前看）的数，然后再从后往前找到第一个比之前找到的那个数大的，交换，然后再反转非递增序列
// 如果已经是非递增的，那么反转整个数组
func nextPermutationCore(nums []int) []int{
	if len(nums)==1{
		return nums
	}
	i := len(nums) - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := len(nums) - 1
		for j >= i && nums[j] <= nums[i] {
			j--
		}
		nums[i],nums[j]=nums[j],nums[i]
	}
	// 反转
	for j:=i+1;j<(len(nums)-1+i+1+1)/2;j++{
		nums[j],nums[len(nums)+i-j]=nums[len(nums)+i-j],nums[j]
	}
	return nums
}
