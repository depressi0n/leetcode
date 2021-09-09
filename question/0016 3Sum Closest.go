package question

import (
	"math"
	"sort"
)

// 题目描述：给定一个包括n 个整数的数组nums和 一个目标值target。
// 找出nums中的三个整数，使得它们的和与target最接近。
// 返回这三个数的和。假定每组输入只存在唯一答案。

func threeSumClosest(nums []int, target int) int {
	return threeSumClosestCore(nums, target)
}

// 先排序，固定一个值，和上一题类似，双指针的思路，并考虑往中间靠拢，离target越近越好
func threeSumClosestCore(nums []int, target int) int {
	compare := func(a, b int) bool {
		if a < 0 {
			a = -a
		}
		if b < 0 {
			b = -b
		}
		if a < b {
			return true
		} else {
			return false
		}
	}
	sort.Ints(nums)
	diff := math.MaxInt32
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}
		left, right := i+1, len(nums)-1

		for left < right {
			// 是否需要更新
			if compare(nums[i]+nums[left]+nums[right]-target, diff) {
				diff = nums[i] + nums[left] + nums[right] - target
				if diff == 0 {
					return target
				}
			}
			// 如果此时的值比target大，那么将right往左去重移动
			if nums[i]+nums[left]+nums[right] > target {
				newRight := right - 1
				for left < newRight && nums[newRight] == nums[right] {
					newRight--
				}
				right = newRight
			}else{// 如果此时的值比target小，那么将left往右去重移动
				newLeft:=left+1
				for newLeft<right && nums[newLeft] == nums[left]{
					newLeft++
				}
				left=newLeft
			}
		}

	}
	return target + diff
}
