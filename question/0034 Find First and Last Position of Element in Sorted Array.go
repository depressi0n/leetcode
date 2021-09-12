package question

// 题目描述：给定一个按照升序排列的整数数组 nums，和一个目标值 target。
// 找出给定目标值在数组中的开始位置和结束位置。
// 如果数组中不存在目标值 target，返回[-1, -1]。
func searchRange(nums []int, target int) []int {
	return searchRangeCore(nums,target)
}

// 主要思想：有序数组中通过二分来查找左右边界
func searchRangeCore(nums []int, target int) []int {
	// 寻找第一个大于等于target的下标，寻找第一个大于target的下标
	// 二者判断条件一致
	binarySearch:=func(nums []int,target int, lower bool)int{
		left,right:=0,len(nums)-1
		res:=len(nums)
		for left<=right{
			mid:=(right-left)/2+left
			if nums[mid] >target || (lower && nums[mid] >=target){
				right=mid -1
				res=mid
			}else{
				left=mid+1
			}
		}
		return res
	}
	leftIdx:=binarySearch(nums,target,true)
	rightIdx:=binarySearch(nums,target,false)-1
	if leftIdx<=rightIdx && rightIdx<len(nums) && nums[leftIdx] == target && nums[rightIdx]==target{
		return []int{leftIdx,rightIdx}
	}
	return []int{-1,-1}
}