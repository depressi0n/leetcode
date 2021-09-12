package question

//题目描述：给你一个有序数组 nums ，请你 原地 删除重复出现的元素，使每个元素 只出现一次 ，返回删除后数组的新长度。
//不要使用额外的数组空间，你必须在 原地 修改输入数组 并在使用 O(1) 额外空间的条件下完成。


func removeDuplicates(nums []int) int {
	return removeDuplicatesCore(nums)
}

// 使用一个下标记录没有重复的元素，因为原下标必然会大于等于，所以不会出现错误
func removeDuplicatesCore(nums []int) int {
	k:=0
	for i:=0;i<len(nums);i++{
		nums[k]=nums[i]
		k++
		for i+1<len(nums) && nums[i+1]==nums[i]{
			i++
		}
	}
	return k
}