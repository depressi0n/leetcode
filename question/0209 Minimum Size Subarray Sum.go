package question

// 滑动窗口，双指针
func minSubArrayLen(s int, nums []int) int {
	if nums[0] >= s {
		return 1
	}
	start, end := 0, 1
	cur := nums[0]
	res := len(nums) + 1
	for {
		for cur < s && end < len(nums) {
			cur += nums[end]
			end++
		}
		if cur < s {
			break
		}
		if end-start < res {
			res = end - start
		}
		cur -= nums[start]
		start++
	}
	if res == len(nums)+1 {
		return 0
	}
	return res
}

//TODO：前缀和+二分查找
