package question

func WinnerSquareGame(n int) bool {
	// 移除规则按照一个数的平方移除
	// 同样可以按照动态规划去解，但是使用递推应该会更方便
	// dp[i]表示先手的人是否能赢得比赛
	dp := make([]bool, n+1)
	// 初始化
	for i := 1; i*i < n+1; i++ {
		dp[i*i] = true
		for j := i*i + 1; j < (i+1)*(i+1) && j < n+1; j++ {
			dp[j] = false
			for k := 1; k <= i; k++ {
				if !dp[j-k*k] {
					dp[j] = true
					break
				}
			}
		}
	}
	return dp[n]
}
