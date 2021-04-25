package question

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	step := 1
	right := nums[0]
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
