package question

import "strings"

// 题目描述:给定两个以字符串形式表示的非负整数 num1 和 num2，返回 num1 和 num2 的乘积，它们的乘积也表示为字符串形式。

func multiply(num1 string, num2 string) string {
	return multiplyCore(num1, num2)
}

// 模拟乘法
func multiplyCore(num1 string, num2 string) string {
	res := make([]int, len(num1)+len(num2))
	for i := len(num1) - 1; i >= 0; i-- {
		carry := 0
		for j := len(num2) - 1; j >= 0; j-- {
			res[len(num1)-i-1+len(num2)-j-1] += int(num1[i]-'0')*int(num2[j]-'0') + carry
		}
	}
	build := strings.Builder{}
	carry := 0
	for i := 0; i < len(res); i++ {
		res[i] = res[i] + carry
		carry = res[i] / 10
		res[i] = res[i] % 10
	}
	right := len(res) - 1
	for ; right > 0 && res[right] == 0; right-- {
		// pass
	}
	for right >= 0 {
		build.WriteByte(byte('0' + res[right]))
		right--
	}
	return build.String()
}

// TODO：FTT，一个string看成一个多项式，使用傅立叶变换加速
