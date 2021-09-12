package question

import "sort"

// 题目描述：给定一个数组candidates和一个目标数target，找出candidates中所有可以使数字和为target的组合。
//candidates中的每个数字在每个组合中只能使用一次。
//注意：解集不能包含重复的组合。

func combinationSum2(candidates []int, target int) [][]int {
	return combinationSum2Core(candidates, target)
}

// 0-1 背包问题
func combinationSum2Core(candidates []int, target int) [][]int {
	// 处理candidates
	sort.Ints(candidates)
	cnt := make(map[int]int)
	for _, num := range candidates {
		cnt[num]++
	}
	// 去重
	k := 0
	for i := 0; i < len(candidates); i++ {
		if candidates[i] != candidates[k] {
			k++
			candidates[k] = candidates[i]
		}
	}
	candidates = candidates[:k+1]
	res := make([][]int, 0, 150)
	// 最简单的思路是通过dfs来直接寻找
	var dfs func(cur int, sum int, path []int)
	dfs = func(cur int, sum int, path []int) {
		if cur == len(candidates) {
			if sum == target {
				t := make([]int, len(path))
				copy(t, path)
				res = append(res, t)
			}
			return
		}
		if sum > target {
			return
		}
		// 选择当前值继续遍历
		if cnt[candidates[cur]] > 0 {
			path = append(path, candidates[cur])
			cnt[candidates[cur]]--
			dfs(cur, sum+candidates[cur], path)
			path = path[:len(path)-1]
			cnt[candidates[cur]]++
		}
		// 不选择当前值
		dfs(cur+1, sum, path)
	}
	dfs(0, 0, []int{})
	return res
}
