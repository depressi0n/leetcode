package question

// 给你一个字符串 s，请你将 s 分割成一些子串，使每个子串都是 回文串 。返回 s 所有可能的分割方案。
//回文串 是正着读和反着读都一样的字符串。
func partition0131(s string) [][]string {
	return partition0131Core2(s)
}

// 主要思想：要返回所有可能，这意味着每一遍都要尝试才行
// 借助动态规划的思想，dp[i][j] 定义为s[i:j]是否是一个回文串
// 然后在dp的帮助下，进行DFS遍历每一种可能
func partition0131Core1(s string) [][]string {
	//动态规划可解，从长度为1到最长，应该可以用并查集，而且会更加简单一点
	//马拉车能不能用进去？不太确定，应该会有问题
	if len(s) == 0 {
		return [][]string{}
	}
	res := make([][]string, 0)
	dp := make([][]bool, len(s)+1) //dp[i][j]表示s[i:j]是否是回文
	for i := 0; i < len(s)+1; i++ {
		dp[i] = make([]bool, len(s)+1)
		dp[i][i] = true //空串
		if i < len(s) {
			dp[i][i+1] = true
		}
	}
	dp[len(s)][len(s)] = true
	for d := 2; d < len(s)+1; d++ {
		for i := 0; i < len(s)+1; i++ {
			if i+d-1 < len(s) && s[i] == s[i+d-1] {
				if i == i+d-2 || dp[i+1][i+d-1] {
					dp[i][i+d] = true
				}
			}
		}
	}
	//对dp进行dfs遍历求解
	var dfs func(int, []string)
	dfs = func(i int, stack []string) {
		if i == len(s) {
			tmp := make([]string, len(stack))
			for i := 0; i < len(stack); i++ {
				tmp[i] = stack[i]
			}
			res = append(res, tmp)
		}
		for j := i + 1; j <= len(s); j++ {
			if dp[i][j] {
				stack = append(stack, s[i:j])
				dfs(j, stack)
				stack = stack[:len(stack)-1]
			}
		}
	}
	for i := 1; i <= len(s); i++ {
		if dp[0][i] {
			dfs(i, []string{s[0:i]})
		}
	}
	//fmt.Println(res)
	return res
}

// 可以将上面的动态规划换成记忆化搜索
func partition0131Core2(s string) [][]string {
	f := make([][]int8, len(s))
	for i := 0; i < len(s); i++ {
		f[i] = make([]int8, len(s))
	}
	// -1表示不是回文串，0表示尚未搜索，1表示是回文串
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		// 空字符串
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
		return f[i][j]
	}
	res := make([][]string, 0, 100)
	splits := make([]string, 0, 100)
	var dfs func(cur int)
	dfs = func(cur int) {
		if cur == len(s) {
			res = append(res, append([]string(nil), splits...))
			return
		}
		for i := cur; i < len(s); i++ {
			if isPalindrome(cur, i) > 0 {
				splits = append(splits, s[cur:i+1])
				dfs(i + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return res
}
