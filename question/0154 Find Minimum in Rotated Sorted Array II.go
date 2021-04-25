package question

func findMin154_1(nums []int) int {
	if len(nums) == 1 || nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	start, end := 0, len(nums)-1
	for start <= end {
		if end == start {
			return nums[start]
		}
		if end == start+1 {
			if nums[end] > nums[start] {
				return nums[start]
			}
			return nums[end]
		}
		mid := start + (end-start)/2
		// 判断都加强一下
		pre := mid
		for pre >= start && nums[pre] == nums[mid] {
			pre--
		}
		post := mid
		for post <= end && nums[post] == nums[mid] {
			post++
		}
		if post > end {
			end = pre + 1
			continue
		}
		if pre < start {
			start = post - 1
			continue
		}
		if nums[mid] > nums[post] {
			return nums[post]
		}
		if nums[pre] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			start = post
		} else {
			end = pre
		}
	}
	return -1
}

//官方的优化版本，用图的方式解答会更直观
// 确实可能回退到O(n)
func findMin154_2(nums []int) int {
	low, high := 0, len(nums)-1
	for low < high {
		pivot := low + (high-low)/2
		if nums[pivot] < nums[high] { // 说明右半边不用考虑
			high = pivot
		} else if nums[pivot] > nums[high] { //说明左半边不用考虑
			low = pivot + 1
		} else { //缩小范围
			high--
		}
	}
	return nums[low]
}
