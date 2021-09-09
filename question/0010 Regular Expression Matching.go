package question

// 题目描述：
// 给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。
//
//'.' 匹配任意单个字符
//'*' 匹配零个或多个前面的那一个元素
//所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

func isMatch(s string, p string) bool {
	return isMatchCore(s,p)
}

// 核心思想：动态规划，遇到*可以原地不动，也可以继续使用
// 状态定义：dp[i][j] 表示s[:i]与p[:j]是否匹配
// if p[j-1] !=*
// dp[i+1][j] = dp[i][j-1] + 1 if s[i] == p[j-1] || p[j-1] == '.' else false
// if p[j-1]=='*'  then dp[i+1][j] =  dp[i][j] || dp[i+1][j-2]  if s[i] == p[j-2] 【0 or多次】else dp[i+1][j-2]
func isMatchCore(s string,p string)bool{
	dp:=make([][]bool,len(s)+1)
	for i:=0;i<len(s)+1;i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	// 初始化
	dp[0][0]=true
	for j:=1;j<len(p)+1;j++{
		if p[j-1] == '*'{
			dp[0][j]=dp[0][j-2]
		}
	}
	// 状态转移
	for i:=1;i<len(s)+1;i++{
		for j:=1;j<len(p)+1;j++{
			// 直接匹配上了的情况
			if s[i-1] == p[j-1] || p[j-1]=='.'{
				dp[i][j]=dp[i-1][j-1]
			}else if p[j-1] == '*'{
				// 无论是否match(s[i-1],p[j-2])，都可以不用这个*号
				dp[i][j]=dp[i][j-2]
				// 但是一旦match(s[i-1],p[j-2]) == true，那么此时*可以给前面用很多次，当然在这里用一次
				if s[i-1] == p[j-2]  || p[j-2] == '.' {
					// 用一次 或 多次
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

