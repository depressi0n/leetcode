package question

import "strings"

func isPalindrome1(s string) bool {
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
