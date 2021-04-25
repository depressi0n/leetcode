package question

import "fmt"

//TODO 很慢，不知道是为啥？看看别人的解析
func minCut(s string) int {
	if len(s) == 0 {
		return 0
	}
	//res:=make([][]string,0)
	dp := make([][]int, len(s)+1) //dp[i][j]表示s[i:j]最少需要几次划分
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]int, len(s)+1)
	}
	for d := 2; d < len(s)+1; d++ {
		for i := 0; i < len(s)-d+1; i++ {

			if dp[i+1][i+d-1] == 0 && s[i] == s[i+d-1] {
				continue
			}
			dp[i][i+d] = dp[i+1][i+d-1] + 2
			for k := i; k <= i+d; k++ {
				if dp[i][k]+dp[k][i+d]+1 < dp[i][i+d] {
					dp[i][i+d] = dp[i][k] + dp[k][i+d] + 1
				}
			}
		}
	}
	fmt.Println(dp[0][len(s)])
	return dp[0][len(s)]
}
