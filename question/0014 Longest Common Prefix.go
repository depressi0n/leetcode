package question

// 题目描述：编写一个函数来查找字符串数组中的最长公共前缀。
//如果不存在公共前缀，返回空字符串 ""。

func longestCommonPrefix(strs []string) string {
	return longestCommonPrefixCore(strs)
}

func longestCommonPrefixCore(strs []string) string {
	if len(strs)==0{
		return ""
	}else if len(strs) == 1{
		return strs[0]
	}
	var longestComminPrefixOfTwoStr func(a,b string)string
	longestComminPrefixOfTwoStr = func(a,b string)string {
		index:=0
		for index<len(a) && index<len(b) && a[index] == b[index]{
			index++
		}
		return a[:index]
	}
	common:=longestComminPrefixOfTwoStr(strs[0],strs[1])
	for i:=2;i<len(strs);i++{
		common=longestComminPrefixOfTwoStr(common,strs[i])
		if common == ""{
			return ""
		}
	}
	return common
}