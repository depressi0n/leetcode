package question

import "math"

// 给你一个非负整数 x ，计算并返回x的 平方根 。
//由于返回类型是整数，结果只保留 整数部分 ，小数部分将被 舍去 。
//注意：不允许使用任何内置指数函数和算符，例如 pow(x, 0.5) 或者 x ** 0.5 。
//
func mySqrt(x int) int {
	return mySqrtCore2(x)
}

// 暴力尝试 => 2分查找
func mySqrtCore1(x int) int {
	i := 1
	for x/i >= i {
		i++
	}
	return i - 1
}
func mySqrtCore2(x int) int {
	start, end := 0, x
	res := -1
	for start <= end {
		mid := start + (end-start)/2
		if mid*mid <= x {
			res = mid
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return res
}

//牛顿迭代
func mySqrtCore3(x int) int {
	if x == 0 { //特例
		return 0
	}
	//变成浮点数
	C, x0 := float64(x), float64(x)
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 { //查看距离，如果几乎不减少了，就认为达成条件可以跳出去
			break
		}
		x0 = xi
	}
	return int(x0)
}

// 袖珍计算器
// sqrt(x)=x^(1/2)=e^(1/2 * ln(x))
func mySqrtCore4(x int) int {
	if x == 0 {
		return 0
	}
	res := int(math.Exp(0.5 * math.Log(float64(x))))
	if (res+1)*(res+1) <= x {
		return res + 1
	}
	return res
}
