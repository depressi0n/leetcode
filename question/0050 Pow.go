package question

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n）。
// -100.0 < x < 100.0
//-2^31 <= n <= 2^31-1
//-10^4 <= x^n <= 10^4
func myPow(x float64, n int) float64 {
	if n > 0 {
		return myPowCore(x, n)
	}
	return 1.0 / myPowCore(x, -n)

}
func myPowCore(x float64, n int) float64 {
	// 指数幂算法
	res := 1.0
	t := x
	for n > 0 {
		if n&0x01 == 1 {
			res *= t
		}
		n >>= 1
		t *= t
	}
	return res
}
