package question

// 题目描述：给定一个无重复元素的正整数数组candidates和一个正整数target，找出candidates中所有可以使数字和为目标数target的唯一组合。
//candidates中的数字可以无限制重复被选取。如果至少一个所选数字数量不同，则两种组合是唯一的。
//对于给定的输入，保证和为target 的唯一组合数少于 150 个。

func combinationSum(candidates []int, target int) [][]int {
	return combinationSumCore1(candidates, target)
}

// 无重复元素中找组合，可以无限制选取，即背包问题，但不能有重复结果
// 用DFS直接遍历求结果即可
func combinationSumCore1(candidates []int, target int) [][]int {
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
		path=append(path,candidates[cur])
		dfs(cur,sum+candidates[cur],path)
		path=path[:len(path)-1]
		// 不再选择当前值
		dfs(cur+1,sum,path)
	}
	dfs(0,0,[]int{})
	return res
}

// TODO：既然是对某个目标值，可以考虑先排好序，根据排序后的结果去剪枝，效果应该更好