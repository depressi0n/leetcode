package question

// 若nums[x]<nums[x+1],则说明右边一定有一个峰顶出现，如果在x+1的右边要么比他大，要么比他小
// 如果比他大，则说明可以继续走
// 如果比他小，则说明肯定有峰值在这边，而最边上可以看作负无穷，说明肯定有
// 同理，左边大的话也是这种情况。
func findPeakElement162_1(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	if len(nums) == 2 {
		if nums[0] < nums[1] {
			return 1
		} else {
			return 0
		}
	}
	index := len(nums) / 2
	if nums[index] > nums[index-1] && nums[index] > nums[index+1] {
		return index
	} else {
		res := findPeakElement162_1(nums[:index+1])
		if res != index {
			return res
		}
		return index + findPeakElement162_1(nums[index:])
	}
}

//迭代版本
func findPeakElement162_2(nums []int) int {
	start, end := 0, len(nums)-1
	for start < end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			end = mid
		} else {
			start = mid + 1
		}
	}
	return start
}
