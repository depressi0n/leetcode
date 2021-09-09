package question

import "sort"

// 题目描述：给你一个由 n 个整数组成的数组nums ，和一个目标值 target 。
// 请你找出并返回满足下述全部条件且不重复的四元组[nums[a], nums[b], nums[c], nums[d]] ：
// 0 <= a, b, c, d< n
// a、b、c 和 d 互不相同
// nums[a] + nums[b] + nums[c] + nums[d] == target
// 你可以按 任意顺序 返回答案


func fourSum(nums []int, target int) [][]int {
	return fourSumCore(nums,target)
}

// 固定两个数，思路同三数之和
func fourSumCore(nums []int, target int) [][]int {
	if len(nums)<4{
		return nil
	}
	res:=make([][]int,0,200)
	sort.Ints(nums)
	for i:=0;i<len(nums);i++{
		for j:=i+1;j<len(nums);j++{
			left,right:=j+1,len(nums)-1
			for left<right{
				if nums[i]+nums[j]+nums[left]+nums[right] == target{
					res=append(res,[]int{nums[i],nums[j],nums[left],nums[right]})
					newRight:=right-1
					for left<newRight && nums[newRight] == nums[right]{
						newRight--
					}
					right=newRight
					newLeft:=left+1
					for newLeft<right && nums[newLeft] == nums[left]{
						newLeft++
					}
					left=newLeft
				}else if nums[i]+nums[j]+nums[left]+nums[right] > target{
					newRight:=right-1
					for left<newRight && nums[newRight] == nums[right]{
						newRight--
					}
					right=newRight
				}else{
					newLeft:=left+1
					for newLeft<right && nums[newLeft] == nums[left]{
						newLeft++
					}
					left=newLeft
				}
			}
			for j+1<len(nums) && nums[j+1] == nums[j]{
				j++
			}
		}
		for i+1<len(nums) && nums[i+1] == nums[i]{
			i++
		}
	}
	return res
}