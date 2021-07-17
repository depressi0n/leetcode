package question

func ProfitableSchemes(n int, minProfit int, group []int, profit []int) int {
	// 目标，获得超过minProfit的收益，投入的人数不超过n个人，不同项目的投入人数和收益不一致
	// 策略：像一个0-1背包问题，直接DFS会超时
	// 加一点剪枝策略：（1）如果在某个情况下已经满足了，可以不再进入dfs，直接通过计数就可以得出来这一个分支下可能的情况【即有点像动态规划】
	//res:=0
	//var dfs func(cur int,total int,remain int)
	//dfs= func(cur int, total int, remain int) {
	//	if cur == len(profit){
	//		if total>=minProfit{
	//			res++
	//		}
	//		return
	//	}
	//	// 可以选并且选择的情况
	//	if remain>=group[cur]{
	//		dfs(cur+1,total+profit[cur],remain-group[cur])
	//	}
	//	// 可以选，但是没选 或者 不能选的情况
	//	dfs(cur+1,total,remain)
	//}
	//dfs(0,0,n)
	//return res

	// 优化思路（1）：动态规划
	// dp[i][j][k]表示nums[0:i]能付出j+1的代价，所能收获的收益至少为k的方案数
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}
	//const mod int = 1e9 + 7
	//ng := len(group)
	//dp := make([][][]int, ng+1)
	//for i := range dp {
	//	dp[i] = make([][]int, n+1)
	//	for j := range dp[i] {
	//		dp[i][j] = make([]int, minProfit+1)
	//	}
	//}
	//dp[0][0][0] = 1
	//// 转移方程
	//// dp[i+1][j][k] = dp[i][j-group[i][k-profit[i]] + dp[i][j][k]
	//for i, members := range group {
	//	earn := profit[i]
	//	for j := 0; j <= n; j++ {
	//		for k := 0; k <= minProfit; k++ {
	//			if j < members {
	//				dp[i+1][j][k] = dp[i][j][k]
	//			} else {
	//				dp[i+1][j][k] = (dp[i][j][k] + dp[i][j-members][max(0, k-earn)]) % mod
	//			}
	//		}
	//	}
	//}
	//cnt := 0
	//for _, d := range dp[ng] {
	//	cnt = (cnt + d[minProfit]) % mod
	//}
	//return cnt

	// 进一步优化成二维
	const mod int = 1e9 + 7
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, minProfit+1)
		dp[i][0] = 1
	}
	for i, members := range group {
		earn := profit[i]
		for j := n; j >= members; j-- {
			for k := minProfit; k >= 0; k-- {
				dp[j][k] = (dp[j][k] + dp[j-members][max(0, k-earn)]) % mod
			}
		}
	}
	return dp[n][minProfit]
}
