package question

import "fmt"

// 格雷编码是一个二进制数字系统，在该系统中，两个连续的数值仅有一个位数的差异。
//给定一个代表编码总位数的非负整数 n，打印其格雷编码序列。即使有多个不同答案，你也只需要返回其中一种。
//格雷编码序列必须以 0 开头。

func grayCode(n int) []int {
	return grayCodeCore(n)
}
// 有点像一个动态规划的思路
// 第n+1层顺序可以由第n层推到出来：
// 在第n层的基础上，加上前缀"0"
// 然后将倒序第n层后，加上前"1"
// 将上述两个集合并起来即可
// 初始化时取n=1  {0,1}
// n=0的情况特判
func grayCodeCore(n int) []int {
	if n == 0 {
		return []int{0}
	}
	res := []int{0, 1}
	for i := 1; i < n; i++ {
		for j := len(res) - 1; j >= 0; j-- {
			res = append(res, res[j]|(1<<i))
		}
	}
	fmt.Println(res)
	return res
}
