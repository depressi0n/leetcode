package question

// 给你一个字符串 s ，字符串的「能量」定义为：只包含一种字符的最长非空子字符串的长度。请返回字符串的能量。

func maxPower(s string) int {
	return maxPowerCore(s)
}

// 主要思想：滑动窗口
func maxPowerCore(s string) int {
	res := 1
	left, right := 0, 1
	for right < len(s) {
		for right < len(s) && s[right] == s[left] {
			right++
		}
		res = max(res, right-left)
		left, right = right, right+1
	}
	return res
}
