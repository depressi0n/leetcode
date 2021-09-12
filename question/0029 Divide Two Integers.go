package question

import "math"

// 题目描述：给定两个整数，被除数dividend和除数divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。
//返回被除数dividend除以除数divisor得到的商。
//整数除法的结果应当截去（truncate）其小数部分，例如：truncate(8.345) = 8 以及 truncate(-2.7335) = -2

func divide(dividend int, divisor int) int {
	return divideCore(dividend, divisor)
}

// 不允许使用乘除法，则使用加减法+移位运算代替
// 优先考虑特殊情况：被除数为0，除数为+/-1
func divideCore(dividend int, divisor int) int {
	var help func(a int, b int) int
	help = func(a int, b int) int {
		// 两个正数
		if a < b {
			return 0
		}
		cnt := 1
		// 不断翻倍，cnt也不断翻倍
		t := b
		for t<<1 <= a {
			cnt <<= 1
			t <<= 1
		}
		// 每次减去尽可能多的结果，此时a-t在[0,2b)之间，结果或0或1
		return cnt + help(a-t, b)

	}
	// 被除数为0
	if dividend == 0 {
		return 0
	}
	// 除数为1
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		}
		return math.MaxInt32
	}
	// 符号位
	sign := 0
	if dividend < 0 {
		sign = sign ^ 1
		dividend = -dividend
	}
	if divisor < 0 {
		sign = sign ^ 1
		divisor = -divisor
	}
	if sign == 1 {
		sign = -1
	} else {
		sign = 1
	}
	res := help(dividend, divisor)
	if sign == 1 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	}
	return -res
}
