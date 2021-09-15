package question

// 给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。

func groupAnagrams0049(strs []string) [][]string {
	return groupAnagrams0049Core(strs)
}
func groupAnagrams0049Core(strs []string) [][]string {
	res:=make([][]string,0,len(strs))
	m2index:=map[int]int{}
	for i := 0; i < len(strs); i++ {
		m:=[26]int{}
		for j := 0; j < len(strs[i]); j++ {
			m[int(strs[i][j]-'a')]++
		}
		t:=0
		for j := 0; j < 26; j++ {
			t=t*100+m[j]
		}
		if _,ok:=m2index[t];!ok{
			res=append(res,[]string{strs[i]})
			m2index[t]=len(res)-1
		}else{
			res[m2index[t]]=append(res[m2index[t]],strs[i])
		}
	}
	return res
}