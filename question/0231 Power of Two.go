package question

func IsPowerOfTwo(n int) bool {
	return isPowerOfTwo1(n)
}
func isPowerOfTwo1(n int) bool {
	if n <= 0 {
		return false
	}
	// 二分查找
	start, end := 0, 31
	for start < end {
		mid := (end-start)/2 + start
		if 1<<mid < n {
			start = mid + 1
		} else if 1<<mid == n {
			return true
		} else {
			end = mid
		}
	}
	if 1<<start == n {
		return true
	}
	return false
}
func isPowerOfTwo2(n int) bool {
	// n & (n-1) 可以将n的二进制表示中最低位1移除
	// 判断是否是0即可
	return n > 0 && n&(n-1) == 0
	// n & (-n) 可以直接获取n的二进制表示中最低位的1
	return n > 0 && n&(-n) == n
}
