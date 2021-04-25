package question

import "fmt"

func partition1(s string) [][]string {
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
	fmt.Println(res)
	return res
}
