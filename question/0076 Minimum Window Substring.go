package question

import (
	"math"
)

// 给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。
func minWindow(s string, t string) string {
	return minWindowCore3(s, t)
}

// 这个方法只适用于t串是无重复的情况
// 如果要修改和这个方法，以适用于有重复的情况
// 需要对重复字符进行特殊处理：以后的每次出现都是之前一次出现的indexs去掉最前的一个
// 这样的话复杂度就上去了
func minWindowCore(s string, t string) string {
	indexs := make([][]int, len(t))
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				indexs[j] = append(indexs[j], i)
			}
		}
	}
	//特殊处理 start
	repeatIndex := make(map[byte][]int) //存放重复字母的下标
	for i := 0; i < len(t); i++ {
		_, ok := repeatIndex[t[i]]
		if !ok {
			repeatIndex[t[i]] = []int{i}
		} else {
			repeatIndex[t[i]] = append(repeatIndex[t[i]], i)
		}
	}
	for _, v := range repeatIndex {
		if len(v) > 1 {
			for i := 1; i < len(v); i++ {
				indexs[v[i]] = indexs[v[i]][i:]
			}
		}
	}
	//特殊处理 end
	for i := 0; i < len(t); i++ {
		if len(indexs[i]) == 0 {
			return ""
		}
	}
	//肯定是递增排序的,并且不可能重复
	min, max := 0, 0
	for i := 1; i < len(t); i++ {
		//找到当前的极值
		if indexs[i][0] > indexs[max][0] { //不需要等于是因为不会出现重复的
			max = i
		}
		if indexs[i][0] < indexs[min][0] { //不需要等于是因为不会出现重复的
			min = i
		}
	}
	rmin, rmax := indexs[min][0], indexs[max][0]
	//移动最小值以更新差值
	for {
		if len(indexs[min]) == len(repeatIndex[t[min]]) {
			break
		}
		//去掉这个最小值
		//indexs[min]=indexs[min][1:]
		// 特殊处理 start
		for i := 0; i < len(repeatIndex[t[min]]); i++ {
			indexs[repeatIndex[t[min]][i]] = indexs[repeatIndex[t[min]][i]][1:]
		}

		// 特殊处理 end
		// 找到当前最大值
		//找到当前的最小值
		min, max = 0, 0
		for i := 1; i < len(t); i++ {
			if indexs[i][0] < indexs[min][0] { //不需要等于是因为不会出现重复的
				min = i
			}
			if indexs[i][0] > indexs[max][0] { //不需要等于是因为不会出现重复的
				max = i
			}
		}
		if indexs[max][0]-indexs[min][0] <= rmax-rmin {
			rmin, rmax = indexs[min][0], indexs[max][0]
		}
	}
	return s[rmin : rmax+1]
}

// 上面这个方法本质上就是滑动窗口的思想
// 用双指针法替代上面的过程
func minWindowCore2(s string, t string) string {
	left, right, minLen := 0, 0, math.MaxInt32
	needSet := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		needSet[t[i]]++
	}
	windowsSet := make(map[byte]int)
	valid := 0 //记录当前窗口中已达成目标的数量，这样就不用每次都去比较map
	start := 0 //记录最短串的开始处
	for right < len(s) {
		if _, ok := needSet[s[right]]; ok { //是需要的字符才会增加
			windowsSet[s[right]]++
			if windowsSet[s[right]] == needSet[s[right]] { //这里用等于的好处是刚好达成，和下面配套更新
				valid++
			}
		}
		right++ //最后的长度和这里是相关的

		for valid == len(needSet) { //开始收缩窗口，可能收缩多次
			if right-left < minLen {
				start = left
				minLen = right - left
			}
			if _, ok := needSet[s[left]]; ok { //只有是需要的字符才减
				windowsSet[s[left]]--
				if windowsSet[s[left]] < needSet[s[left]] { // 只有少于目标值时才会将valid减少
					valid--
				}
			}
			left++
		}
	}
	if minLen == math.MaxInt32 {
		return ""
	}
	return s[start : start+minLen]
}

// 考虑以下情况：
// （1）t中字符可能重复
// （2）复杂度尽可能底，双指针实现
func minWindowCore3(s string, t string) string {
	if len(s)<len(t){
		return ""
	}
	start, end := 0, len(t)
	currentSet := make(map[byte]int)
	needSet := make(map[byte]int)
	// 初始化
	for i := 0; i < len(t); i++ {
		needSet[t[i]]++
	}
	cnt := 0
	for i := start; i < len(t); i++ {
		if _, ok := needSet[s[i]]; ok {
			currentSet[s[i]]++
			if currentSet[s[i]] == needSet[s[i]] {
				cnt++
			}
		}
	}
	minStart, minLen := len(s)+1, len(s)+1
	if cnt == len(needSet){
		return s[start:end]
	}
	for end < len(s) {
		if n, ok := needSet[s[end]]; ok {
			currentSet[s[end]]++
			if currentSet[s[end]] == n {
				cnt++
			}
		}
		end++
		for cnt == len(needSet) {
			//if minLen < len(s) {
			//	fmt.Println(s[minStart : minStart+minLen])
			//}
			if minLen > end-start {
				minStart, minLen = start, end-start
			}
			//fmt.Println(s[minStart : minStart+minLen])
			// 去掉第一个在needSet中元素
			_, ok := needSet[s[start]]
			if ok {
				currentSet[s[start]]--
				if currentSet[s[start]] == needSet[s[start]]-1 {
					cnt--
				}
			}
			start++
		}
	}
	if minLen == len(s)+1 {
		return ""
	}
	return s[minStart : minStart+minLen]
}
