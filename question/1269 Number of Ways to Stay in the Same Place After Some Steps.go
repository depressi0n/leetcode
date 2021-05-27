package question

func NumWays(steps int, arrLen int) int {
	// 思路肯定没错，但是这么做肯定会超时
	//res:=0
	//var ways  func(remainSteps int,curLeft int)
	//ways= func(remainSteps int,curLeft int){
	//	if remainSteps == 0 && curLeft==0{
	//		res++
	//	}
	//	if remainSteps<curLeft{
	//		// 说明不可能回答原点
	//		return
	//	}
	//	if curLeft>=arrLen{
	//		// 越界了
	//		return
	//	}
	//	// 停留在原地
	//	ways(remainSteps-1,curLeft)
	//	// 向左再走一步
	//	ways(remainSteps-1,curLeft+1)
	//	// 向右走一步
	//	if curLeft>0{
	//		ways(remainSteps-1,curLeft-1)
	//	}
	//	return
	//}
	//ways(steps,0)
	//return res

	// 另一条思路，但是会很麻烦，从排列组合的角度出发
	// 用数学的方法处理
	// 1+C(steps,2)*Catan(1)+C(steps,2)*Catan(2)+...+C(steps,2*k)*Catan(k+1)+【这里开始就不好计算了】
	// 结合这种方式和下面的动态规划，是不是有一些有意思的事情发生

	// 动态规划
	// dp[i][j] 表示在第i步操作下，指针位于j的方案数
	// i [0,steps]
	// j [0,min(steps,arrLen-1)]
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	mod := 1000000007
	distances := min(steps, arrLen-1)
	dp := make([][]int, steps+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, distances+1)
	}
	// 初始化
	dp[0][0] = 1
	// 状态转移方程，TODO 可以进行状态压缩
	// dp[i][j] = dp[i-1][j-1] + dp[i-1][j] + dp[i-1][j+1]
	for i := 1; i < len(dp); i++ {
		for j := 0; j < distances+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j-1 >= 0 {
				dp[i][j] = (dp[i][j] + dp[i-1][j-1]) % mod
			}
			if j+1 <= distances {
				dp[i][j] = (dp[i][j] + dp[i-1][j+1]) % mod
			}
		}
	}
	return dp[steps][0]
}
