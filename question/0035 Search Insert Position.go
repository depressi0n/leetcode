package question

// 题目描述：给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。
//
//请必须使用时间复杂度为 O(log n) 的算法。

func searchInsert(nums []int, target int) int {
	return searchInsertCore(nums,target)
}

// 如果不存在，应该被插入的位置第一个大于它的位置
func searchInsertCore(nums []int, target int) int {
	left,right:=0,len(nums)-1
	for left<=right{
		mid:=(right+left)/2
		if nums[mid]== target{
			return mid
		}else if nums[mid]<target{
			left=mid+1
		}else{
			right=mid-1
		}
	}
	return left
}