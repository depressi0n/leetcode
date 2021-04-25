package question

import "math"

func wordBreak139_1(s string, wordDict []string) bool {
	// 第一想法就是用map记录单词
	// 记录单词列表中最短和最长，然后从开始往后匹配
	min, max := math.MaxInt32, 0
	m := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = struct{}{}
		if len(wordDict[i]) < min {
			min = len(wordDict[i])
		}
		if len(wordDict[i]) > max {
			max = len(wordDict[i])
		}
	}
	//开始查找
	find := func(index, length int) bool {
		if _, ok := m[s[index:index+length]]; ok {
			return true
		}
		return false
	}
	var dfs func(int) bool
	dfs = func(index int) bool {
		if index == len(s) {
			return true
		}
		for l := min; l <= max; l++ {
			if index+l <= len(s) && find(index, l) {
				if dfs(index + l) {
					return true
				}
			}
		}
		return false
	}
	return dfs(0)
}
func wordBreak139_2(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
