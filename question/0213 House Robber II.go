package question

//
func rob(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}
	max := func(x, y int) int {
		if x < y {
			return y
		}
		return x
	}
	robt := func(nums []int) int {
		if len(nums) == 0 {
			return 0
		}
		if len(nums) == 1 {
			return nums[0]
		}
		first := nums[0]
		second := nums[0]
		if nums[0] < nums[1] {
			second = nums[1]
		}
		for i := 2; i < len(nums); i++ {
			first, second = second, max(first+nums[i], second)
		}
		return second
	}
	// 第一个和最后一个只能选择一个，所以要么答案在前面部分，要么答案在后面部分
	return max(robt(nums[0:len(nums)-1]), robt(nums[1:]))
}

//TODO:证明这样的结论
