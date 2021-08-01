package question

func Answer1540(s string, t string, k int) bool {
	return canConvertString(s, t, k)
}
func canConvertString(s string, t string, k int) bool {
	// k步以内，把s变成t
	// 第i次可以选一个没有被选过的index，shift i 次，round
	// 思路：首先长度不一致，直接返回false
	if len(s) != len(t) {
		return false
	}
	nums := make([]int, 26)
	// 计算二者各个index上的差值，负数则+26
	for i := 0; i < len(s); i++ {
		// error n:=(int(t[i]-s[i])+26)%26 因为里面是一个uint8，一旦负数则溢出了
		n := int(t[i]-s[i]+26) % 26
		if n == 0 {
			continue
		}
		if nums[n]*26+n > k {
			return false
		}
		nums[n]++
	}
	return true
}
