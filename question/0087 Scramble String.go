package question

func isScramble(s1 string, s2 string) bool {
	// 主要思想
	// 1、如果两个字符串是通过这种方式来的，那么必定可以分成两段分别得到的，只是分割线的位置不一样
	// 2、根据上面的思想，可以想到用dp做
	// 3、由于需要确定一个分割线的位置，所以dp至少需要两个维度，一个维度给s1，一个维度给s2
	// 4、因为长度不一样肯定不会是通过扰乱得到的，所以再加上一个维度去表示长度
	// 5、dp[i][j][k] 表示 s1[i:i+k]和s2[j:j+k]能否通过扰乱得到
	// 6、第三维的k是核心内容，最后的结果是dp[0][0][len(s1)]  为了方便调试，将长度维度放到第一层
	if len(s1) != len(s2) {
		return false
	}
	dp := make([][][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([][]bool, len(s1))
		for j := 0; j < len(s1); j++ {
			dp[i][j] = make([]bool, len(s1))
		}
	}
	//初始化，长度为1的情况
	for i := 0; i < len(s1); i++ {
		for j := 0; j < len(s1); j++ {
			if s1[i] == s2[j] {
				dp[1][i][j] = true
			}
		}
	}
	//状态转移方程
	//dp[i][j][k]=(dp[i][j][1]&&dp[i+1][j+1][k-1]) || (dp[i][j][2]&&dp[i+2][j+2][k-2]) || ... || (dp[i][j][k-1]&&dp[i+k-1][j+k-1][1])
	// 需要一定程度的判断，因为长度肯定是不够的
	// 边界情况的处理
	for k := 2; k < len(s1)+1; k++ {
		for i := 0; i < len(s1); i++ {
			if i+k > len(s1) {
				break
			}
			flag := false
			for j := 0; j < len(s1); j++ {
				if j+k > len(s1) {
					flag = true
					break
				}
				for l := 1; l < k; l++ {
					dp[k][i][j] = dp[k][i][j] || (dp[l][i][j] && dp[k-l][i+l][j+l]) || (dp[l][i][j+k-l] && dp[k-l][i+l][j])
				}
			}
			if flag {
				continue
			}
		}
	}
	//结果
	return dp[len(s1)][0][0]
}
