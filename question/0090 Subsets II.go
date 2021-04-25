package question

import "sort"

func subsetsWithDup(nums []int) [][]int {
	//首先分析题意，nums可能含有重复元素，然后需要划分子集，但要求子集不能重复
	//直觉上的思路是选与不选的问题
	//但是这么做带来的问题是会出现重复的子集
	//解决问题的直观想法是去比较结果，然后不接受重复的就行，但是这么做复杂度会比较高
	//思路：空集、一个元素的去重、两个元素也要去重
	// 有点像背包了，所以复杂肯定是上去了，这样的话就可以先把nums排序
	// 然后按照子集的大小进行排序
	// 做dfs选择的时候跳过相同值的，只选择一个即可
	res := [][]int{{}}
	if len(nums) == 0 {
		return res
	}
	var dfsIn90 func(int, int, []int)
	dfsIn90 = func(cur int, size int, s []int) {
		if len(s) == size {
			tmp := make([]int, size)
			for i := 0; i < size; i++ {
				tmp[i] = s[i]
			}
			res = append(res, tmp)
			return
		}
		if cur >= len(nums) {
			return
		}
		s = append(s, nums[cur])
		dfsIn90(cur+1, size, s) //choose nums[cur]

		s = s[:len(s)-1]
		for cur+1 < len(nums) && nums[cur+1] == nums[cur] {
			cur++
		}
		dfsIn90(cur+1, size, s) // do not choose nums[cur]
	}
	sort.Ints(nums)
	for size := 1; size <= len(nums); size++ {
		cur := 0
		dfsIn90(cur, size, []int{})
	}
	return res
}
