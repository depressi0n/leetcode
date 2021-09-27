package question

// 给定三个字符串s1、s2、s3，请你帮忙验证s3是否是由s1和s2 交错 组成的。
//两个字符串 s 和 t 交错 的定义与过程如下，其中每个字符串都会被分割成若干 非空 子字符串：
//s = s1 + s2 + ... + sn
//t = t1 + t2 + ... + tm
//|n - m| <= 1
//交错 是 s1 + t1 + s2 + t2 + s3 + t3 + ... 或者 t1 + s1 + t2 + s2 + t3 + s3 + ...
//提示：a + b 意味着字符串 a 和 b 连接。

//TODO:滚动数组
func isInterleave(s1 string, s2 string, s3 string) bool {
	return isInterleaveCore(s1, s2, s3)
}
// 主要思想，动态规划
// dp[i][j] 表示s1[:i] 和 s2[:j] 能否按照上述规则得到s3[:i+j-1]
func isInterleaveCore(s1 string, s2 string, s3 string) bool {
	// 长度不等，直接判否
	if len(s1)+len(s2) != len(s3) {
		return false
	}
	// dp[i][j]表示s1[0,i]和s2[0,j]是否是按照这种规则得到s3[0,i+j]
	dp := make([][]bool, len(s1)+1)
	for i := 0; i < len(s1)+1; i++ {
		dp[i] = make([]bool, len(s2)+1)
	}

	// 初始化，dp[0][0]=true 意味着能从""和""得到""
	dp[0][0] = true
	// 如果出现了子串的情况，则一直往后均为true，断开后均为false
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
