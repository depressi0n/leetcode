package question

func searchII(nums []int, target int) bool {
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
