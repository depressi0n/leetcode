package question

// 假设你正在爬楼梯。需要 n 阶你才能到达楼顶。
//每次你可以爬 1 或 2 个台阶。你有多少种不同的方法可以爬到楼顶呢？
//注意：给定 n 是一个正整数。
func climbStairs(n int) int {
	return climbStairsCore(n)
}
// 最简单的动态规划思想：反向思考，你可以怎么下楼梯，最后回到下一楼层
func climbStairsCore(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2
	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}
