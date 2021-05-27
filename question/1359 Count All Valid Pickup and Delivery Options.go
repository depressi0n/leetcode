package question

func CountOrders(n int) int {
	// 有点类似于括号匹配，但每个括号都是有标号的，本质上是一个排列组合问题
	// 把所有左括号先排好 即一个全排列 保证随后的序列肯定不重复
	// 开始插入右括号，插入的顺序有讲究
	// 最先插入的是最后一个左括号对应的右括号，然后插入倒数第二个左括号对应的右括号，以此类推
	// 最后的值应当是 A_n^n * π_1^n (2n-1) = n! * (2 * n! - 1)
	//  最后的结果要mod 10^9+7
	res := int64(1)
	for i := 1; i <= n; i++ {
		res = res * int64(i*(2*i-1)) % 1000000007
	}
	return int(res)
}
