package question

// 题目描述：给定一个字符串(s) 和一个字符模式(p) ，实现一个支持'?'和'*'的通配符匹配。
//'?' 可以匹配任何单个字符。
//'*' 可以匹配任意字符串（包括空字符串）。
//两个字符串完全匹配才算匹配成功。
//说明:
//s可能为空，且只包含从a-z的小写字母。
//p可能为空，且只包含从a-z的小写字母，以及字符?和*。

func isMatch0044(s string, p string) bool {
	return isMatch0044Core2(s, p)
}

// 使用动态规划
// dp[i][j]表示s[:i],p[:j]是否能匹配上
// 初始化时 dp[0][0]=true,dp[1][0]=false,dp[1][j]
// 状态转移
// if p[j]=='*' then dp[i][j]=dp[i-1][j-1]
// if p[j]!='*' then dp[i][j] = s[i-1]==p[j-1] || p[j-1]='?'
func isMatch0044Core1(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	// 初始化
	dp[0][0] = true // ["",""] 相同
	for j := 1; j < len(p)+1 && p[j-1] == '*'; j++ {
		dp[0][j] = true
	}
	for i := 1; i < len(s)+1; i++ {
		for j := 1; j < len(p)+1; j++ {
			if p[j-1]!='*'{
				if s[i-1]==p[j-1] || p[j-1]=='?'{
					dp[i][j]=dp[i-1][j-1]
				}
			}else{
				dp[i][j]=dp[i][j-1]
				for k:=i-1;k>=0;k--{
					dp[i][j]=dp[i][j] || dp[k][j-1]
					if dp[i][j]{
						break
					}
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}

// 贪心算法
func isMatch0044Core2(s string, p string) bool {
	for len(s)>0 && len(p)>0  && p[len(p)-1]!='*'{
		if s[len(s)-1]==p[len(p)-1]||p[len(p)-1]=='?'{
			s=s[:len(s)-1]
			p=p[:len(p)-1]
		}else{
			return false
		}
	}
	if len(p)==0{
		return len(s)==0
	}
	curS,curP:=0,0
	preS,preP:=-1,-1
	for curS<len(s) && curP<len(p){
		if p[curP] == '*'{
			curP++
			preS,preP=curS,curP
		}else if s[curS]==p[curP] || p[curP]=='?'{
			curS++
			curP++
		}else if preS!=-1 && preS+1<len(s){
			// 回退
			preS++
			curS,curP=preS,preP
		}else{
			return false
		}
	}
	for curP<len(p) && p[curP]=='*'{
		curP++
	}
	return curP==len(p)
}