package question

func numTrees(n int) int {
	return numTreesCore(n)
}
// TODO:Catana数
// 主要思想：
// 首先明确当n==1时，就只有1种，当n==2时，有2种，当n==3时，有左1右1，左0右2，左2右0，共5种
// 使用动态规划 dp[i]+=dp[j-1]*dp[i-j] 1<=j<=i
func numTreesCore(n int) int {
	if n<=1{
		return 1
	}
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = 2 * dp[0] * dp[i-1] //第一个和最后一个作为root
		for j := 2; j < i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
