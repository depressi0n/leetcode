package question

// 题目描述：给定一个字符串s和一些 长度相同 的单词words 。找出 s 中恰好可以由words 中所有单词串联形成的子串的起始位置。
// 注意子串要与words 中的单词完全匹配，中间不能有其他字符 ，但不需要考虑words中单词串联的顺序。

func findSubstring(s string, words []string) []int {
	return findSubstringCore(s, words)
}

// (1)words中所有字符串长度相等
// (2)words中所有字符串都要出现，可能会有多个相同的情况
// 以上两点，借助一个长度length和一个map来表征
// 遍历s的每一个位置，按照s[i:i+length]查询map是否出现过
func findSubstringCore(s string, words []string) []int {
	if len(s)<len(words)*len(words[0]){
		return []int{}
	}
	res := make([]int, 0, len(s)-len(words)*len(words[0]))
	// preparation
	length := len(words[0])
	m := make(map[string]int, len(words))
	for i := 0; i < len(words); i++ {
		m[words[i]]++
	}
	for i := 0; i <= len(s)-len(words)*len(words[0]); i++ {
		start := i
		t := make(map[string]int, len(words))
		for start < i+len(words)*length {
			t[s[start:start+length]]++
			start += length
		}
		flag := true
		for subStr, cnt := range t {
			if m[subStr] != cnt {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
	//fmt.Println(res)
	return res
}

// TODO：多起点滑动窗口？
// 主要思想：维护多个滑动窗口(i=0,...,d-1)，除非第一个单词离开了窗口，否则窗口内其他单词未受到影响，每次滑动d