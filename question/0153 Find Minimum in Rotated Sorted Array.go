package question

func findMin153_1(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// 无论怎么变，二分查找的策略同样是有效的
	var bFind func(int, int) int
	bFind = func(start int, end int) int {
		if end-start == 2 {
			if nums[start] < nums[start+1] {
				return nums[start]
			}
			return nums[start+1]
		}
		if nums[start] < nums[end-1] { //单调的
			return nums[start]
		}
		mid := start + (end-start)/2
		if nums[start] < nums[mid] { //最小值肯定出现在后半部分,带着mid跑
			return bFind(mid, end)
		}
		return bFind(start, mid+1) // 带着mid跑
	}
	return bFind(0, len(nums))
}

func findMin153_2(nums []int) int {
	if len(nums) == 1 || nums[len(nums)-1] > nums[0] {
		return nums[0]
	}
	start, end := 0, len(nums)-1
	for start <= end {
		mid := start + (end-start)/2
		if nums[mid] > nums[mid+1] {
			return nums[mid+1]
		}
		if nums[mid-1] > nums[mid] {
			return nums[mid]
		}
		if nums[mid] > nums[0] {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}
