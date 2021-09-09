package question

import "strings"

// 题目描述：给定一个仅包含数字2-9的字符串，返回所有它能表示的字母组合。答案可以按 任意顺序 返回。
//给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
// 2 -> abc
// 3 -> def
// 4 -> ghi
// 5 -> jkl
// 6 -> mno
// 7 -> pqrs
// 8 -> tuv
// 9 -> wxyz
func letterCombinations(digits string) []string {
	return letterCombinationsCore(digits)
}

// 首先硬编码，然后dfs遍历
func letterCombinationsCore(digits string) []string {
	if len(digits) == 0{
		return nil
	}
	res := make([]string,0,len(digits)*4)
	m := map[byte][]byte{
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
	t := strings.Builder{}
	var dfs func(cur int, pre []byte)
	dfs = func(cur int, pre []byte) {
		if cur == len(digits) {
			t.Reset()
			t.Write(pre)
			res = append(res, t.String())
			return
		}
		for i := 0; i < len(m[digits[cur]]); i++ {
			pre = append(pre, m[digits[cur]][i])
			dfs(cur + 1,pre)
			pre=pre[:len(pre)-1]
		}
	}
	dfs(0,[]byte{})
	return res
}
