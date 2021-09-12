package question

// 题目描述：给你两个字符串haystack 和 needle ，
// 请你在 haystack 字符串中找出 needle 字符串出现的第一个位置（下标从 0 开始）。如果不存在，则返回-1 。

func strStr(haystack string, needle string) int {
	return strStrCore1(haystack,needle)
}

// 暴力匹配
func strStrCore1(haystack string, needle string) int {
	if needle==""{
		return 0
	}
	for i:=0;i<len(haystack);i++{
		j:=0
		for ;i+j<len(haystack) && j<len(needle) && haystack[i+j]==needle[j];j++{
			// pass
		}
		if j== len(needle){
			return i
		}
	}
	return -1
}

// KMP算法，核心是理解前缀函数（即next数组）
// next[i] 表示 s[0:i+1]中最长的相等的真前缀和真后缀的长度，如果不存在则为0
// if s[i]=s[next[i-1]], then next[i]=next[i-1]+1
// 举个🌰
// aabaaab
// a,next[0]=0，a没有真前缀和真后缀
// aa,next[1]=1，s[0:1]=s[1:2]
// aab,next[2]=0
// aaba,next[3]=1,s[0:1]=s[3:4]
// aabaa,next[4]=2,s[0:2]=s[3:5]
// aabaaa,next[5]=2,s[0:2]=s[4:6]
// aabaaab,next[6]=3,s[0:3]=s[4:7]
func strStrCore2(haystack string, needle string) int {
	if needle==""{
		return 0
	}
	// 求next数组
	getNextArr:=func(s string)[]int{
		res:=make([]int,len(s))
		i,j:=1,0 // 长度为0
		for ;i<len(s);i++{
			for j>0 && s[i]!=s[j]{
				j=res[j-1]
			}
			if s[i]==s[j]{
				j++
			}
			res[i]=j
		}
		return res
	}
	next:=getNextArr(needle)
	i,j:=0,0
	for ;i<len(haystack);i++{
		for j>0 && haystack[i] !=needle[j]{
			j=next[j-1]
		}
		if haystack[i]==needle[j]{
			j++
		}
		if j==len(needle){
			return i-j+1
		}
	}
	return -1
}