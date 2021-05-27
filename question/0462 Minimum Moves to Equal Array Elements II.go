package question

import (
	"fmt"
	"sort"
)

func MinMoves2(nums []int) int {
	// 保证所有的数字都相同，考虑中位数
	// f(x)=|x-nums[0]|+|x-nums[1]|+|x-nums[2]|+...+|x-nums[len(nums)-1]|
	// 求f(x)最小值，在中位数的时候取到
	// 题意转换为求中位数  可以用快排的思想取求中位数
	if len(nums) < 2 {
		return 0
	}
	var findK func(nums []int, k int) int
	findK = func(nums []int, k int) int {
		pivot := nums[0]
		i, j := 0, len(nums)-1
		for {
			for ; j > i && nums[j] > pivot; j-- {

			}
			if j == i {
				nums[i] = pivot
				break
			}
			nums[i] = nums[j]
			i++
			for ; i < j && nums[i] <= pivot; i++ {

			}
			if i == j {
				nums[i] = pivot
				break
			}
			nums[j] = nums[i]
			j--
		}
		if i+1 < k {
			return findK(nums[i+1:], k-i-1)
		} else if i+1 == k {
			return pivot
		} else {
			return findK(nums[:i], k)
		}
	}
	midV := findK(nums, (len(nums)+1)/2)
	fmt.Printf("midV=%v\n", midV)
	res := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] < midV {
			res += midV - nums[i]
		} else {
			res += nums[i] - midV
		}
	}
	return res
}

// TODO: 接受平衡的思想
func minMoves2_2(nums []int) int {
	sort.Ints(nums)
	s, e := 0, len(nums)-1
	steps := 0
	for s < e {
		steps += nums[e] - nums[s]
		s++
		e--
	}
	return steps
}
