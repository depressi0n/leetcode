package question

//给定两个字符串s和 p，找到s中所有p的异位词的子串，返回这些子串的起始索引。不考虑答案输出的顺序。
//异位词 指由相同字母重排列形成的字符串（包括相同的字符串）。
func findAnagrams(s string, p string) []int {
	return findAnagramsCore(s, p)
}

// 主要思想：滑动窗口的思想，首先预处理统计p的字符，然后初始化窗口去处理每个窗口的字符的情况，并且在这个过程中记录满足条件的字符的个数matched
// 然后判定是否符合字符个数要求，并移动窗口，如果左侧元素移除窗口后不再足够，则将matched-1即可，右侧元素移入窗口，同样地，维护matched
func findAnagramsCore(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	// 预处理p
	//处理s
	pm := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		pm[p[i]]++
	}
	wm := make(map[byte]int)
	matched := 0
	for i := 0; i < len(p); i++ {
		wm[s[i]]++
		if wm[s[i]] == pm[s[i]] {
			matched++
		}
	}
	left, right := 0, len(p)
	res := make([]int, 0, len(s))
	for right <= len(s) {
		if matched == len(pm) {
			res = append(res, left)
		}
		if right==len(s){
			break
		}
		// 去掉左边
		if wm[s[left]] == pm[s[left]] {
			matched--
		}
		wm[s[left]]--
		// 加上右边
		wm[s[right]]++
		if wm[s[right]] == pm[s[right]] {
			matched++
		}
		left++
		right++
	}
	return res
}
