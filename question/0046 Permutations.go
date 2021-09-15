package question

//题目描述：给定一个不含重复数字的数组 nums ，返回其 所有可能的全排列 。你可以 按任意顺序 返回答案。

func permute(nums []int) [][]int {
	return permuteCore(nums)
}

func permuteCore(nums []int) [][]int {
	checkEqualSlice := func(t, nums []int) bool {
		i := 0
		for ; i < len(nums); i++ {
			if t[i] != nums[i] {
				break
			}
		}
		return i == len(nums)
	}
	res := make([][]int, 0, len(nums))
	//  保证了nums中整数互不相同
	t := make([]int, len(nums))
	copy(t, nums)
	res = append(res, t)
	for {
		tmp := make([]int, len(nums))
		copy(tmp, res[len(res)-1])
		next := nextPermutationCore0046(tmp)
		if !checkEqualSlice(next, nums) {
			res = append(res, next)
		} else {
			break
		}
	}
	return res
}
func nextPermutationCore0046(nums []int) []int {
	// 首先找到从后往前第一个非递减的
	if len(nums)==1{
		return nums
	}
	i := len(nums) - 1
	for ; i > 0 && nums[i-1] >= nums[i]; i-- {
		// pass
	}
	i--
	// 如果整个序列从左往右是非递增序列，直接交换整个数组
	if i>=0{
		// 非递增序列中找到比这个值大的，可以使用二分法加快
		j := len(nums) - 1
		for ; j >= i && nums[j] <= nums[i]; j-- {
			// pass
		}
		// 交换
		nums[i], nums[j] = nums[j], nums[i]
	}
	// 反转
	for k := i + 1; k < (len(nums)+i+1)/2; k++ {
		nums[k], nums[len(nums)+i-k] = nums[len(nums)+i-k], nums[k]
	}
	return nums
}
