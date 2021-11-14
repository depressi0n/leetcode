package question

// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,4,4,5,6,7] 在变化后可能得到：
//若旋转 4 次，则可以得到 [4,5,6,7,0,1,4]
//若旋转 7 次，则可以得到 [0,1,4,4,5,6,7]
//注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
//给你一个可能存在 重复 元素值的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
func findMin0154(nums []int) int {
	return findMin0154Core2(nums)
}

// 与上一题的区别在于可能存在重复元素，此时旋转之后的结果会出现一些偏差
// 主要思想：在缩小范围的时候加强判定，如果范围缩小到1或2，同样是平凡的结果
// 如果范围长度大于2，则将与中间值相同的元素全部去掉，再缩小范围
func findMin0154Core(nums []int) int {
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

// 官方的优化版本，用图的方式解答会更直观
// 确实可能回退到O(n)
// 主要思想是：如果中间值小于尾则后半段必然有序，如果中间值大于尾部值则前半段必然有序，如果相等，则只能将尾部一个一个往前推进
// 考虑数组中的最后一个元素x，在最小值右侧的元素，它们的值一定都小于等于x
// 而在最小值左侧的元素，它们的值一定都大于等于x
// 因此，我们可以根据这一条性质，通过二分查找的方法找出最小值。
func findMin0154Core2(nums []int) int {
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
