package question

import "strings"

// 题目描述：给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串 的长度。
func lengthOfLongestSubstring(s string) int {
	return lengthOfLongestSubstring2(s)
}

// 双指针+map，双指针指向当前的无重复字符的子串，map记录出现过的字符的位置
func lengthOfLongestSubstring1(s string) int {
	if len(s)<=1{
		return len(s)
	}
	res:=1
	left,right:=0,1
	m:=map[byte]int{s[0]:0}
	for right<len(s){
		if _,ok:=m[s[right]];!ok{
			m[s[right]]=right
			right++
		}else {
			// 当前的长度
			if right-left>res{
				res=right-left
			}
			nextLeft := m[s[right]]
			for ; left < nextLeft; left++ {
				delete(m, s[left])
			}
			m[s[left]]=right
			left++
			right++
		}
	}
	if right-left>res{
		res=right-left
	}
	return res
}

// 借助库函数strings.Index寻找，效率上更快
// 查看当前位置在之前是否出现过，如果没有出现过，则看情况更新end
// 如果出现过，则更新start和end
func lengthOfLongestSubstring2(s string) int {
	start,end := 0,0
	for i := 0; i < len(s); i++ {
		index := strings.Index(s[start:i],string(s[i]))
		if index==-1{
			if i+1>end{
				end=i+1
			}
		}else{
			start+=index+1
			end+=index+1
		}
	}
	return end-start
}