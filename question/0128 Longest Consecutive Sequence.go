package question

import (
	"sort"
)

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//请你设计并实现时间复杂度为O(n) 的算法解决此问题。
func longestConsecutive(nums []int) int {
	return longestConsecutiveCore2(nums)
}

// 直观：先排序，然后根据排序结果来获得最长连续序列
// 复杂度为O(nlogn)
func longestConsecutiveCore1(nums []int) int {
	sort.Ints(nums)
	res, cur := 1, 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1]+1 {
			cur++
			if cur > res {
				res = cur
			}
		} else {
			cur = 1
		}
	}
	return res
}

// 优化方向：使用一个map来映射数据是否存在
// 对数组中的元素进行遍历的过程中，如果元素存在，则往两边扩展，边扩展边删除，并维护一个最大值，最后得到的结果就是最大。如果元素不存在则跳过去（说明被之前的删除了）

func longestConsecutiveCore2(nums []int) int {
	m := make(map[int]struct{}, len(nums))
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	if len(nums) <1{
		return 0
	}
	res := 1
	for i := 0; i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			left, right := nums[i]-1, nums[i]+1
			for _, exist := m[right]; exist;  {
				delete(m, right)
				right++
				_, exist = m[right]
			}
			for _, exist := m[left]; exist;  {
				delete(m, left)
				left--
				_, exist = m[left]
			}
			delete(m, nums[i])
			res = max(res, right-left-1)
		}
	}
	return res
}
func longestConsecutiveCore3(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	m := make(map[int]struct{})
	for i := 0; i < len(nums); i++ {
		m[nums[i]] = struct{}{}
	}
	max := 1
	for k, _ := range m {
		start := k - 1
		end := k + 1
		for {
			if _, ok := m[start]; ok {
				delete(m, start)
				start--
			} else {
				break
			}
		}

		for {
			if _, ok := m[end]; ok {
				delete(m, end)
				end++
			} else {
				break
			}
		}
		delete(m, k)
		if end-start-1 > max {
			max = end - start - 1
		}
	}
	return max
}
// TODO：考虑用并查集实现
// 并查集的状态定义为left和right
// 向并查集中插入一个元素时，可能导致合并，这时候需要更新left和right
// 这样每个元素必将属于某一个状态
// 遍历过程中，维护一个最长长度，遍历完所有的元素之后，维护的最长长度就是最后结果