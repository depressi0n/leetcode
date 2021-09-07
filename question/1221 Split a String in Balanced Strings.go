package question
// 题目描述：在一个 平衡字符串 中，'L' 和 'R' 字符的数量是相同的。
//给你一个平衡字符串 s，请你将它分割成尽可能多的平衡字符串。
//注意：分割得到的每个字符串都必须是平衡字符串。
//返回可以通过分割得到的平衡字符串的 最大数量 。
func balancedStringSplit(s string) int {
	return balancedStringSplitCore(s)
}
// 贪心即可
func balancedStringSplitCore(s string) int {
	res:=0
	cnt:=0
	cur:=0
	for cur<len(s){
		if s[cur]=='R' {
			cnt++
		}else{
			cnt--
		}
		if cnt == 0{
			res++
		}
		cur++
	}
	return res
}