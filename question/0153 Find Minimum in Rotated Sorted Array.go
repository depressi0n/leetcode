package question

// 已知一个长度为 n 的数组，预先按照升序排列，经由 1 到 n 次 旋转 后，得到输入数组。例如，原数组 nums = [0,1,2,4,5,6,7] 在变化后可能得到：
// 若旋转 4 次，则可以得到 [4,5,6,7,0,1,2]
// 若旋转 7 次，则可以得到 [0,1,2,4,5,6,7]
// 注意，数组 [a[0], a[1], a[2], ..., a[n-1]] 旋转一次 的结果为数组 [a[n-1], a[0], a[1], a[2], ..., a[n-2]] 。
// 给你一个元素值 互不相同 的数组 nums ，它原来是一个升序排列的数组，并按上述情形进行了多次旋转。请你找出并返回数组中的 最小元素 。
func findMin0153(nums []int) int {
	return findMin0153Core1(nums)
}
// 主要思想：旋转之后，二分查找的策略依旧有效：这是因为数组仍旧是部分有序的
// 如果查找的长度变成2，则已经是基本情况，比较大小返回结果
// 如果查找长度大于2，则首先比较首尾，如果首小于尾，说明是有序的，直接返回
// 如果首大于尾，说明最小值在中间，拿出中间位置的数，进一步缩小范围，如果中间位置上的数比首元素小，则范围为首到中间，但要带着中间走，否则说明结果在中间到尾部
func findMin0153Core1(nums []int) int {
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
// 直接使用二分查找，和上面的思想一致，只是代码更加规整
func findMin0153Core2(nums []int) int {
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
