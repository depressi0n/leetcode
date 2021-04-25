package question

func lengthOfLastWord(s string) int {
	res := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && res == 0 {
			continue
		}
		if s[i] == ' ' {
			return res
		}
		res++
	}
	return res
}
