package question

//  一条包含字母 A-Z 的消息通过以下映射进行了 编码 ：
//'A' -> 1
//'B' -> 2
//...
//'Z' -> 26
//要 解码 已编码的消息，所有数字必须基于上述映射的方法，反向映射回字母（可能有多种方法）。例如，"11106" 可以映射为：
//"AAJF" ，将消息分组为 (1 1 10 6)
//"KJF" ，将消息分组为 (11 10 6)
//注意，消息不能分组为 (1 11 06) ，因为 "06" 不能映射为 "F" ，这是由于 "6" 和 "06" 在映射中并不等价。
//给你一个只含数字的 非空 字符串 s ，请计算并返回 解码 方法的 总数 。
//题目数据保证答案肯定是一个 32 位 的整数。
func numDecodings(s string) int {
	return numDecodingsCore(s)
}
// 主要思想，两个数可以绑定，特别是遇到0时必须绑定
// 遍历每一个数，如果可以和前面绑定起来，则绑定起来也计算，否则和前面相等
func numDecodingsCore(s string) int {
	// 特殊情况
	if len(s) == 0 || s[0] == '0' {
		return 0
	}
	// 初始化0的位置
	prev, next := 1, 1
	for i := 1; i < len(s); i++ {
		switch {
		case s[i] == '0' && s[i-1] != '1' && s[i-1] != '2': //非法
			return 0
		case s[i] == '0': //上面case后半部分的补集
			next = prev
		case s[i-1] == '1' || (s[i-1] == '2' && s[i] <= '6'): //二者可以绑定起来
			prev, next = next, prev+next
		default: //一般情况就是前后相等即无法绑定
			prev = next
		}
	}
	return next
}
