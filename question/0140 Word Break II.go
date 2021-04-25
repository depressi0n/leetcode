package question

import "strings"

func wordBreak140_1(s string, wordDict []string) []string {
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

// 必须用记忆化搜索，不然对于最坏情况会空间受限而移除
func wordBreak140_2(s string, wordDict []string) (sentences []string) {
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
