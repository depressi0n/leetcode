package question

import "math"

// 给你一个字符串 s 和一个字符串列表 wordDict 作为字典，判定 s 是否可以由空格拆分为一个或多个在字典中出现的单词。
//说明：拆分时可以重复使用字典中的单词。
func wordBreak0139(s string, wordDict []string) bool {
	return wordBreak0139Core2(s, wordDict)
}

// 朴素想法：使用一个map记录所有的单词，并记录最小长度和最大长度，然后对s从后往前匹配，使用记忆化搜索
func wordBreak0139Core1(s string, wordDict []string) bool {
	// 第一想法就是用map记录单词
	// 记录单词列表中最短和最长，然后从开始往后匹配
	minLength, maxLength := math.MaxInt32, 0
	m := make(map[string]struct{})
	for i := 0; i < len(wordDict); i++ {
		m[wordDict[i]] = struct{}{}
		if len(wordDict[i]) < minLength {
			minLength = len(wordDict[i])
		}
		if len(wordDict[i]) > maxLength {
			maxLength = len(wordDict[i])
		}
	}
	// 记忆化,cache[i] 表示 s[i:] 是否能匹配上
	cache := make([]int, len(s)+1)
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
		for l := minLength; l <= maxLength; l++ {
			if index+l <= len(s) && find(index, l) {
				// 查询cache
				if cache[index+l] == -1 {
					continue
				}
				if cache[index+l] == 1 {
					return true
				} else {
					flag := dfs(index + l)
					if flag {
						cache[index+l] = 1
						return true
					} else {
						cache[index+l] = -1
					}
				}
			}
		}
		return false
	}
	return dfs(0)
}
// 将记忆化搜索修改动态规划，从后往前遍历
func wordBreak0139Core2(s string, wordDict []string) bool {
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
