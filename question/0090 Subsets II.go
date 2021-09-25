package question

import "sort"

//给你一个整数数组 nums ，其中可能包含重复元素，请你返回该数组所有可能的子集（幂集）。
//解集 不能 包含重复的子集。返回的解集中，子集可以按 任意顺序 排列。

func subsetsWithDup(nums []int) [][]int {
	return subsetsWithDupCore(nums)
}
func subsetsWithDupCore(nums []int) [][]int {
	//首先分析题意，nums可能含有重复元素，然后需要划分子集，但要求子集不能重复
	//直觉上的思路是选与不选的问题
	//但是这么做带来的问题是会出现重复的子集
	//解决问题的直观想法是去比较结果，然后不接受重复的就行，但是这么做复杂度会比较高

	// 最终思路：空集、一个元素的去重、两个元素也要去重
	// 先排序，然后选取第一个元素，每次选取一个不一样的元素
	// 有点像背包了，所以复杂肯定是上去了，这样的话就可以先把nums排序
	// 然后按照子集的大小进行排序
	// 做dfs选择的时候跳过相同值的，只选择一个即可
	res := [][]int{{}}
	if len(nums) == 0 {
		return res
	}
	var dfs func(int, int, []int)
	dfs = func(cur int, size int, s []int) {
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
		dfs(cur+1, size, s) //choose nums[cur]

		s = s[:len(s)-1]
		// 选择过的相同元素不必再选取，因为在之前已经得到过结果
		for cur+1 < len(nums) && nums[cur+1] == nums[cur] {
			cur++
		}
		dfs(cur+1, size, s) // do not choose nums[cur]
	}
	sort.Ints(nums)
	for size := 1; size <= len(nums); size++ {
		cur := 0
		dfs(cur, size, []int{})
	}
	return res
}
// TODO：考虑一下迭代法的实现