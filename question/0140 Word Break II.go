package question

import "strings"

// 给定一个非空字符串 s 和一个包含非空单词列表的字典 wordDict，在字符串中增加空格来构建一个句子，使得句子中所有的单词都在词典中。返回所有这些可能的句子。
// 说明：
// 分隔时可以重复使用字典中的单词。
// 你可以假设字典中没有重复的单词。
func wordBreak0140(s string, wordDict []string) []string {
	return wordBreak0140Core3(s, wordDict)
}

// 核心思想：使用字典树来记录单词，主要处理单词的划分
func wordBreak0140Core1(s string, wordDict []string) []string {
	// 记录一个最长单词长度
	minWordLength, maxWordLength := len(s), 0
	m := make(map[string]int)
	for index, word := range wordDict {
		m[word] = index
		maxWordLength = max(maxWordLength, len(word))
		minWordLength = min(minWordLength, len(word))
	}
	// startWith[i] 表示 以s[i]为结尾的单词可能的开头位置
	startWith := make([][]int, len(s))
	endWith := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		for length := minWordLength; length <= maxWordLength && i+length <= len(s); length++ {
			if _, ok := m[s[i:i+length]]; ok {
				endWith[i+length-1] = append(endWith[i+length-1], i)
				startWith[i] = append(startWith[i], i+length-1)
			}
		}
	}
	// 进行一次动态规划，根据dp的状态可能的结果
	// res[i] 表示 s[i:] 被划分的结果
	res := make([][]string, len(s))
	// 初始化时直接就是最后一个单词的情况
	for i := 0; i < len(endWith[len(s)-1]); i++ {
		res[endWith[len(s)-1][i]] = append(res[endWith[len(s)-1][i]], s[endWith[len(s)-1][i]:])
	}
	for i := len(s) - minWordLength - 1; i >= 0; i-- {
		for j := 0; j < len(endWith[i]); j++ {
			for k := 0; k < len(res[i+1]); k++ {
				res[endWith[i][j]] = append(res[endWith[i][j]], s[endWith[i][j]:i+1]+" "+res[i+1][k])
			}
		}
	}
	return res[0]
}
//
func wordBreak0140Core2(s string, wordDict []string) []string {
	m := make(map[string]int)
	for index, word := range wordDict {
		m[word] = index
	}
	res := make([]string, 0)
	dp := make([][][]int, len(s)+1)
	dp[0] = [][]int{{-1}}
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if len(dp[j]) != 0 {
				index, ok := m[s[j:i]]
				if !ok {
					continue
				}
				for _, h := range dp[j] {
					tmp := make([]int, len(h))
					for k := 0; k < len(h); k++ {
						tmp[k] = h[k]
					}
					tmp = append(tmp, index)
					dp[i] = append(dp[i], tmp)
				}
			}
		}
	}
	for i := 0; i < len(dp[len(s)]); i++ {
		tmp := ""
		for j := 1; j < len(dp[len(s)][i]); j++ {
			tmp += wordDict[dp[len(s)][i][j]]
			if j != len(dp[len(s)][i])-1 {
				tmp += " "
			}
		}
		res = append(res, tmp)
	}
	return res
}

// 必须用记忆化搜索，不然对于最坏情况会空间受限而溢出
func wordBreak0140Core3(s string, wordDict []string) (sentences []string) {
	wordSet := map[string]struct{}{}
	for _, w := range wordDict {
		wordSet[w] = struct{}{}
	}

	n := len(s)
	dp := make([][][]string, n)
	var backtrack func(index int) [][]string
	backtrack = func(index int) [][]string {
		if dp[index] != nil {
			return dp[index]
		}
		wordsList := [][]string{}
		for i := index + 1; i < n; i++ {
			word := s[index:i]
			if _, has := wordSet[word]; has {
				for _, nextWords := range backtrack(i) {
					wordsList = append(wordsList, append([]string{word}, nextWords...))
				}
			}
		}
		word := s[index:]
		if _, has := wordSet[word]; has {
			wordsList = append(wordsList, []string{word})
		}
		dp[index] = wordsList
		return wordsList
	}
	for _, words := range backtrack(0) {
		sentences = append(sentences, strings.Join(words, " "))
	}
	return
}
