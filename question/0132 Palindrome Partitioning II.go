package question

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是回文。
//返回符合要求的 最少分割次数 。

func minCut(s string) int {
	return minCutCore2(s)
}

// 主要思想：只要求最少分割次数，所以尽可能贪心，在DFS的过程中，一次分割的字符串尽可能远
// 另外，根据次数进行剪枝
// 借助动态规划的思想：dp[i][j] 表示s[i:j]最少需要几次划分
// 初始化时对长度为0和长度为1的子串为1
// dp[i][j+1]= min(dp[i][k]+dp[k][j+1]) fo k -> [i+1,j+1]
// 这样做的时间复杂度是O(n^3)，会超时
func minCutCore1(s string) int {
	if len(s) == 0 {
		return 0
	}
	//dp[i][j]表示s[i:j]最少需要几次划分
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s)+1)
		dp[i][i] = 0
		dp[i][i+1] = 0
		for j := i + 2; j < len(dp[i]); j++ {
			dp[i][j] = j - i
		}
	}
	for length := 2; length <= len(s); length++ {
		for start := 0; start+length <= len(s); start++ {
			if s[start] == s[start+length-1] && dp[start+1][start+length-1] == 0 {
				dp[start][start+length] = 0
				continue
			}
			for k := start + 1; k < start+length; k++ {
				dp[start][start+length] = min(dp[start][start+length], dp[start][k]+dp[k][start+length]+1)
			}
		}
	}
	//fmt.Println(dp[0][len(s)])
	return dp[0][len(s)]
}

// 主要思想：用两次动态规划代替上面的一次动态规划，这样做的好处是动态规划会将所有需要的都计算出来，但是记忆化搜索可以通过剪枝的方式不去求那些没用的结果
// 时间复杂度可以降低到n^2
func minCutCore2(s string) int {
	if len(s)==0{
		return 0
	}
	// f[i][j] 表示s[i:j+1]是否是一个回文串
	f := make([][]bool, len(s))
	for i := 0; i < len(s); i++ {
		f[i] = make([]bool, len(s))
		for j := 0; j < len(f[i]); j++ {
			f[i][j] = true
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}
	// res[i]表示s[:i+1]至少要经过几次划分
	res := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		if f[0][i] {
			continue
		}
		res[i] = i + 1
		for j := 0; j < i; j++ {
			if f[j+1][i] && res[j]+1 < res[i] {
				res[i] = res[j] + 1
			}
		}
	}
	return res[len(s)-1]
}
