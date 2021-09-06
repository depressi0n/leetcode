package question

func Answer0413(nums []int) int {
	return numberOfArithmeticSlices(nums)
}
func numberOfArithmeticSlices(nums []int) int {
	// 子数组为一个算术数组的个数
	// 算术数组定义：长度至少为3，连续两个数字的差相等
	// 思路：长度判别，双指针法
	if len(nums) < 3 {
		return 0
	}
	res := 0
	left, right := 0, 2                // 起始为2个元素，这样才有差值
	diff := nums[right-1] - nums[left] //差值
	// 首先是找到第一个组，然后右指针尽可能右移
	for right < len(nums) {
		for right <len(nums) && nums[right] == nums[right-1]+diff {
			right++
		}
		for right-left >= 3 {
			res+=right-left-2
			left++
		}
		// 到这里，长度要归为2，左边到最后一个元素，右边右移
		left = right - 1
		right = left + 2
		if right>len(nums){
			break
		}
		diff = nums[right-1] - nums[left] //差值
	}
	return res
}
