package question

// 题目描述：整数数组 nums 按升序排列，数组中的值 互不相同 。
// 在传递给函数之前，nums 在预先未知的某个下标 k（0 <= k < nums.length）上进行了 旋转，使数组变为 [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]（下标 从 0 开始 计数）。
// 例如， [0,1,2,4,5,6,7] 在下标 3 处经旋转后可能变为 [4,5,6,7,0,1,2] 。
// 给你 旋转后 的数组 nums 和一个整数 target ，如果 nums 中存在这个目标值 target ，则返回它的下标，否则返回-1。
func search(nums []int, target int) int {
	return searchCore2(nums,target)
}
// 简化一下
func searchCore1(nums []int, target int) int {
	left,right:=0,len(nums)-1
	var mid int
	for left<right{
		mid=(right+left)/2
		if nums[mid] == target{
			return mid
		}
		if nums[mid]<nums[right]{
			// 后半段有序
			if nums[mid]<target && target<=nums[right]{
				left=mid+1
			}else{
				right=mid-1
			}
		}else if nums[left]<nums[mid]{
			// 前半段有序
			if nums[left]<=target && target<nums[mid]{
				right=mid-1
			}else{
				left=mid+1
			}
		}else{
			if nums[left]!=target{
				left++
			}else{
				return left
			}
			if nums[right]!=target{
				right--
			}else{
				return right
			}
		}
	}
	if nums[left]==target{
		return left
	}
	return -1
}

func searchCore2(nums []int, target int) int {
	if len(nums)== 0 || (len(nums)==1 && nums[0]!=target){
		return -1
	}
	if len(nums)==1 && nums[0]==target {
		return 0
	}
	left,right:=0,len(nums)-1
	for left<=right{
		mid:=(left+right)/2
		if nums[mid] == target{
			return mid
		}
		if nums[left] <=nums[mid]{
			if nums[left]<=target && target<nums[mid]{
				right=mid-1
			}else{
				left=mid+1
			}
		}else{
			if nums[mid]<target && target<=nums[right]{
				left=mid+1
			}else{
				right=mid-1
			}
		}
	}
	return -1
}