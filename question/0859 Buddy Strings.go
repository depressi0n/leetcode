package question

//给你两个字符串 s 和 goal ，只要我们可以通过交换 s 中的两个字母得到与 goal 相等的结果，就返回 true ；否则返回 false 。
//交换字母的定义是：取两个下标 i 和 j （下标从 0 开始）且满足 i != j ，接着交换 s[i] 和 s[j] 处的字符。
//例如，在 "abcd" 中交换下标 0 和下标 2 的元素可以生成 "cbad" 。

func buddyStrings(s string, goal string) bool {
	return buddyStringsCore2(s, goal)
}

// 主要思想：首先检查长度，然后考虑是否出现相同字母，从前往后检查并记录，如果恰有两个字母可以交换，并且交换后相同，那么返回true，否则返回false
func buddyStringsCore1(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	m := make(map[byte]struct{})
	same := false
	diff := -1
	i := 0
	for i < len(s) {
		if _, ok := m[s[i]]; ok {
			same = true
		}
		m[s[i]] = struct{}{}
		if s[i] == goal[i] {
			i++
			continue
		}
		if diff == -1 {
			diff = i
			i++
			continue
		} else if s[diff] != goal[i] || s[i]!=goal[diff] {
			return false
		} else {
			i++
			// 保证后面全部相同
			for i < len(s) && s[i] == goal[i] {
				i++
			}
			if i < len(s) {
				return false
			}
			return true
		}
	}
	if diff == -1 {
		return same
	}
	return false
}
func buddyStringsCore2(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}
	if s==goal{
		// 需要检查是否有重复字符出现，即上面的same的作用
		seen := [26]bool{}
		for _, ch := range s {
			if seen[ch-'a'] {
				return true
			}
			seen[ch-'a'] = true
		}
		return false
	}
	// first 和 second 的作用类似于上面的diff
	first, second := -1, -1
	for i := range s {
		if s[i] != goal[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}
	return second != -1 && s[first] == goal[second] && s[second] == goal[first]
}
