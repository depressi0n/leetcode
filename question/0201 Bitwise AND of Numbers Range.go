package question

func rangeBitwiseAnd(m int, n int) int {
	// 根据数字个数
	// 从最低位开始，0/1交替出现，然后每隔2^i交替出现
	// 一旦AND后，只有全部为1的地方不会出错
	//res:=0
	//v := 1
	//cnt := 0
	//for v < n-m+1 {
	//	cnt++
	//	v *= 2
	//}
	//res := ((m >> cnt) & (n >> cnt)) << cnt
	//return res

	for m < n {
		n &= (n - 1)
	}
	return n
}
