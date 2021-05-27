package question

func StrangePrinter(s string) int {
	// 特性一：一次只会打印相同的字母
	// 特性二：每一次转变可以从任意位置开始，然后覆盖之前的结果

	// 反向看这个题目
	// 每一次可以选择将一串连续的相同的字母变成另一串字母
	// 求最少的次数

	//if len(s) == 0 {
	//	return 0
	//}
	//if len(s) == 1 {
	//	return 1
	//}
	//if len(s) == 2 && s[0] == s[1] {
	//	return 1
	//}
	//if len(s) == 2 && s[0] != s[1] {
	//	return 2
	//}
	////  找到与s[0]相同的字母，到i=0必定停止
	//i := len(s) - 1
	//for i >= 0 && s[i] != s[0] {
	//	i--
	//}
	//if i == 0 {
	//	//	说明没有与第一个字母相同的字母
	//	//  这个字母必须单独打印
	//	return 1+ StrangePrinter(s[1:])
	//}else{
	//	return StrangePrinter(s[:i])+ StrangePrinter(s[i+1:])
	//}

	// //TODO：动态规划的思想，为什么没想到状态转移方程
	// dp[i][j] 表示s[i:j]的最少操作数
	// 状态转移方程
	// if s[i] == s[j] then dp[i][j]=dp[i][j-1]
	// if s[i] != s[j] then dp[i][j] = min(dp[i][k]+dp[k+1][j]) i<=k<=j-1
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
	}
	for i := len(s) - 1; i >= 0; i-- {
		dp[i][i] = 1
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i][j-1]
			} else {
				dp[i][j] = len(s)
				for k := i; k < j; k++ {
					dp[i][j] = func(a, b int) int {
						if a < b {
							return a
						}
						return b
					}(dp[i][j], dp[i][k]+dp[k+1][j])
				}
			}
		}
	}
	return dp[0][len(s)-1]
}
