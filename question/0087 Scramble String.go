package question

// 使用下面描述的算法可以扰乱字符串 s 得到字符串 t ：
//如果字符串的长度为 1 ，算法停止
//如果字符串的长度 > 1 ，执行下述步骤：
//在一个随机下标处将字符串分割成两个非空的子字符串。即，如果已知字符串 s ，则可以将其分成两个子字符串 x 和 y ，且满足 s = x + y 。
//随机 决定是要「交换两个子字符串」还是要「保持这两个子字符串的顺序不变」。即，在执行这一步骤之后，s 可能是 s = x + y 或者 s = y + x 。
//在 x 和 y 这两个子字符串上继续从步骤 1 开始递归执行此算法。
//给你两个 长度相等 的字符串 s1 和s2，判断s2是否是s1的扰乱字符串。如果是，返回 true ；否则，返回 false 。
func isScramble(s1 string, s2 string) bool {
	return isScrambleCore(s1, s2)
}
func isScrambleCore(s1 string, s2 string) bool {
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
	// 可以加上一层判断
	m := [26]int{}
	for i := 0; i < len(s1); i++ {
		m[int(s1[i]-'a')]++
		m[int(s2[i]-'a')]--
	}
	for i := 0; i < len(s1); i++ {
		if m[i] != 0 {
			return false
		}
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
		for i := 0; i+k <= len(s1); i++ {
			for j := 0; j+k <= len(s1); j++ {
				for l := 1; l < k; l++ {
					dp[k][i][j] = dp[k][i][j] || (dp[l][i][j] && dp[k-l][i+l][j+l]) || (dp[l][i][j+k-l] && dp[k-l][i+l][j])
				}
			}
		}
	}
	//结果
	return dp[len(s1)][0][0]
}
