package question

import "unicode"

// 我们定义，在以下情况时，单词的大写用法是正确的：
//全部字母都是大写，比如 "USA" 。
//单词中所有字母都不是大写，比如 "leetcode" 。
//如果单词不只含有一个字母，只有首字母大写，比如"Google" 。
//给你一个字符串 word 。如果大写用法正确，返回 true ；否则，返回 false 。

func detectCapitalUse(word string) bool {
	return detectCapitalUseCore2(word)
}

// 主要思想：设置一个flag表示期待的后续字母是大小写还是均可
// flag 为1时表示期望大写，0表示均可，-1表示期望小写
func detectCapitalUseCore1(word string) bool {
	if len(word) == 1 {
		return true
	}
	// 默认首字母期待均可
	flag := 0
	// 检查首字母，如果是小写，意味着只能跟小写
	if 'a' <= word[0] && word[0] <= 'z' {
		flag = -1
	}
	// 否则要根据第二个字母，确定是否全部是大写
	if flag == 0 && 'A' <= word[1] && word[1] <= 'Z' {
		flag = 1
	} else {
		flag = -1
	}
	for i := 1; i < len(word); i++ {
		if flag == 1 && 'a' <= word[i] && word[i] <= 'z' {
			return false
		}
		if flag == -1 && 'A' <= word[i] && word[i] <= 'Z' {
			return false
		}
	}
	return true
}
// 规则一： 如果第一个字母是小写那么第二个字母也必须是小写
// 规则二： 后续所有字母必须和第二个字母一致
func detectCapitalUseCore2(word string) bool {
	if len(word) >= 2 && unicode.IsLower(rune(word[0])) && unicode.IsUpper(rune(word[1])) {
		return false
	}
	for i := 2; i < len(word); i++ {
		if unicode.IsLower(rune(word[i])) != unicode.IsLower(rune(word[1])) {
			return false
		}
	}
	return true
}
