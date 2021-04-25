package question

import "math"

func mySqrt(x int) int {
	i := 1
	for x/i >= i {
		i++
	}
	return i - 1
}
func mySqrt1(x int) int {
	//牛顿迭代
	if x == 0 { //特例
		return 0
	}
	C, x0 := float64(x), float64(x) //变成浮点数
	for {
		xi := 0.5 * (x0 + C/x0)
		if math.Abs(x0-xi) < 1e-7 { //查看距离，如果几乎不减少了，就认为达成条件可以跳出去
			break
		}
		x0 = xi
	}
	return int(x0)
}
