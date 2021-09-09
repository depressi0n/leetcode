package question

import "sort"

// 题目描述：给你一个包含 n 个整数的数组nums，
// 判断nums中是否存在三个元素 a，b，c ，使得a + b + c = 0 ？
// 请你找出所有和为 0 且不重复的三元组。
// 注意：答案中不可以包含重复的三元组

func threeSum(nums []int) [][]int {
	return threeSumCore2(nums)
}

// 暴力枚举，先排序（为了方便去重），三重循环... 会超时
func threeSumCore1(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))
	for i := 0; i < len(nums); {
		for j := i + 1; j < len(nums); {
			for k := j + 1; k < len(nums); {
				if nums[i]+nums[j]+nums[k] == 0 {
					res = append(res, []int{nums[i], nums[j], nums[k]})
				}
				// 去重
				k++
				for ; k < len(nums) && nums[k] == nums[k-1]; k++ {
				}
			}
			// 去重
			j++
			for ; j < len(nums) && nums[j] == nums[j-1]; j++ {
			}
		}
		i++
		for ; i < len(nums) && nums[i] == nums[i-1]; i++ {
		}

	}
	return res
}

// 排序 + 双指针
func threeSumCore2(nums []int) [][]int {
	if len(nums) < 3 {
		return [][]int{}
	}
	sort.Ints(nums)
	res := make([][]int, 0, len(nums))
	for i := 0; i < len(nums);i++ {
		// 跳过枚举过的
		if i>0 && nums[i-1] == nums[i]{
			continue
		}
		target:=-1*nums[i]
		// 双指针
		left,right:=i+1,len(nums)-1
		for ;left<len(nums);left++{
			// 跳过枚举过的
			if  left>i+1 && nums[left-1] == nums[left] {
				continue
			}
			// 如果此时比目标值大，则right左移变小
			for left<right && nums[left]+nums[right]>target{
				right--
			}
			// 如果此时left和right重合了,重合的在第一轮已经遍历过了，此时右边可能会有相同的值
			if left == right{
				break
			}
			if nums[left]+nums[right] == target{
				res=append(res,[]int{nums[i],nums[left],nums[right]})
			}
		}
	}
	return res
}
