package question

//TODO:滚动数组
func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	//if len(s1) == 0 && strings.Compare(s2, s3) == 0 {
	//	return true
	//}
	//if len(s1) == 0 {
	//	return false
	//}
	//if len(s2) == 0 && strings.Compare(s1, s3) == 0 {
	//	return true
	//}
	//if len(s2) == 0 {
	//	return false
	//}
	////第一感觉就是用dp做
	// dp[i][j]表示s1[0,i]和s2[0,j]是否是按照这种规则得到s3[0,i+j]
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	//初始化
	dp[0][0] = true
	for i := 1; i < len(s1)+1 && s1[i-1] == s3[i-1]; i++ {
		dp[i][0] = true
	}
	for i := 1; i < len(s2)+1 && s2[i-1] == s3[i-1]; i++ {
		dp[0][i] = true
	}
	//状态转移
	// 如果s3[i+j-1] == s1[i-1] dp[i][j]=dp[i-1][j]
	// 如果s3[i+j-1] == s2[j-1] dp[i][j]=dp[i][j-1]
	for i := 1; i < len(s1)+1; i++ {
		for j := 1; j < len(s2)+1; j++ {
			if s1[i-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i-1][j]
			}
			if s2[j-1] == s3[i+j-1] {
				dp[i][j] = dp[i][j] || dp[i][j-1]
			}
		}
	}
	return dp[len(s1)][len(s2)]
}
