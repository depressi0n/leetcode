package question

func myPow(x float64, n int) float64 {
	if n > 0 {
		return pow(x, n)
	}
	return 1.0 / pow(x, -n)

}
func pow(x float64, n int) float64 {
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
