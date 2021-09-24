package question
// 已知存在一个按非降序排列的整数数组 nums ，数组中的值不必互不相同。
//
//在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转 ，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。例如， [0,1,2,4,4,4,5,6,6,7] 在下标 5 处经旋转后可能变为 [4,5,6,6,7,0,1,2,4,4] 。
//
//给你 旋转后 的数组 nums 和一个整数 target ，请你编写一个函数来判断给定的目标值是否存在于数组中。如果 nums 中存在这个目标值 target ，则返回 true ，否则返回 false 。

func searchII(nums []int, target int) bool {
	return searchIICore(nums,target)
}
// 主要思想：如果被旋转之后，start和mid、end和mid出现相同的则去掉首或尾，否则就比较大小往中间靠拢，如果start<end则继续，否则返回
// 是否使用start<=end取决于初始值
func searchIICore(nums []int, target int) bool {
	if len(nums) == 0 {
		return false
	}
	start := 0
	end := len(nums) - 1
	mid := 0
	for start <= end {
		mid = (start + end) / 2
		if nums[mid] == target {
			return true
		}
		// 因为数据可能重复，可能start，end，mid都相等的情况会发生，这样就无法判断有序还是无序了
		// 但是如果这种情况发生，就去掉相同的
		if nums[mid] == nums[end] { //出现相同的就去掉相同的
			end--
			continue
		}
		if nums[start] == nums[mid] { //出现相同的就去掉相同的
			start++
			continue
		}
		// 除了上面的情况，左边和右边必定会出现不想等，那么就可以排序了
		if nums[start] < nums[mid] { //左半部分有序
			if nums[start] <= target && target < nums[mid] {
				end = mid - 1
			} else {
				start = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[end] {
				start = mid + 1
			} else {
				end = mid - 1
			}
		}
	}
	return false
}
