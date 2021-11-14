package question

import "strings"

// 给你一个字符串 s ，逐个翻转字符串中的所有 单词 。
// 单词 是由非空格字符组成的字符串。s 中使用至少一个空格将字符串中的 单词 分隔开。
// 请你返回一个翻转 s 中单词顺序并用单个空格相连的字符串。
// 说明：
// 输入字符串 s 可以在前面、后面或者单词间包含多余的空格。
// 翻转后单词间应当仅用一个空格分隔。
// 翻转后的字符串中不应包含额外的空格。
func reverseWords(s string) string {
	return reverseWordsCore2(s)
}

// 主要思想，首先根据空格分割出所有的单词，需要注意的是原始字符串中单词可能被连续空格所分割
// 然后从后往前组合即可，使用一个空格连接
// 下面一种实现是不经过库函数的方式
func reverseWordsCore1(s string) string {
	//	去掉前后空格
	start := 0
	for s[start] == ' ' {
		start++
	}
	end := len(s) - 1
	for s[end] == ' ' {
		end--
	}
	words := make([]string, 0)
	for i := start; i < end+1; i++ {
		if s[i] == ' ' {
			//单词结束
			words = append(words, s[start:i])

			//去掉连续空格
			for s[i] == ' ' {
				i++
			}

			//重置
			start = i
		}
	}
	words = append(words, s[start:end+1])
	res := ""
	for i := len(words) - 1; i >= 0; i-- {
		res += words[i]
		res += " "
	}
	return res[:len(res)-1]
}

// 借助库函数，直接分割，然后去掉空字符串，然后用空格拼接起来
func reverseWordsCore2(s string) string {
	arr := strings.Split(s, " ")
	res := make([]string, 0, len(arr))
	// 去掉空字符串
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) != 0 {
			res = append(res, arr[i])
		}
	}
	return strings.Join(res, " ")
}
