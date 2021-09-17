package question

//给定一个非负整数数组 nums ，你最初位于数组的 第一个下标 。
//数组中的每个元素代表你在该位置可以跳跃的最大长度。
//判断你是否能够到达最后一个下标。

func canJump(nums []int) bool {
	return canJumpCore(nums)
}
func canJumpCore(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	step := 1
	right := nums[0] // 记录当前能达到的最远位置
	if right >= len(nums)-1 {
		return true
	}
	if right == 0 {
		return false
	}
	for i := 1; i <= right && i < len(nums)-1; i++ {
		if nums[i]+i >= len(nums)-1 {
			return true
		}
		if nums[i]+i > right {
			right = nums[i] + i
			step++
		}
	}
	return false
}
