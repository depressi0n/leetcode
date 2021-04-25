package question

//计算尾部的0，实际上看因子中的5，2可以不用看是因为肯定是足够的
func trailingZeroes(n int) int {
	cnt := 0
	for i := 5; i <= n; i += 5 {
		j := i
		for j%5 == 0 {
			j = j / 5
			cnt++
		}
	}
	return cnt
}
func trailingZeroes1(n int) int {
	cnt := 0
	for n > 0 {
		n /= 5
		cnt += n
	}
	return cnt
}
