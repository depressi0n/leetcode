package question

import "strings"

// 给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。
//说明：本题中，我们将空字符串定义为有效的回文串
func isPalindrome0125(s string) bool {
	return isPalindrome0125Core(s)
}
func isPalindrome0125Core(s string) bool {
	if len(s) < 2 {
		return true
	}
	s = strings.ToLower(s)
	i := 0
	j := len(s) - 1
	for i < len(s) {
		if (s[i] < 'a' || s[i] > 'z') && (s[i] < '0' || s[i] > '9') {
			i++
			continue
		}
		if (s[j] < 'a' || s[j] > 'z') && (s[j] < '0' || s[j] > '9') {
			j--
			continue
		}
		if s[i] == s[j] {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}
